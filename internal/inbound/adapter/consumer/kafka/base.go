package kafka

import (
	"encoding/json"
	"ftgo-order/internal/outbound/interface/logger"
	"ftgo-order/pkg/helpers"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const MaxPartition = 50

type KafkaConsumer struct {
	Consumer       *kafka.Consumer
	Topics         []string
	Config         kafka.ConfigMap
	Logger         logger.Logger
	topicPartition []kafka.TopicPartition
}

func (c *KafkaConsumer) ReadConfig() {
	viper.SetConfigName("kafka")
	viper.SetConfigType("yaml")
	projectPath := helpers.ProjectPath()
	viper.AddConfigPath(projectPath + "config")
	if err := viper.ReadInConfig(); err != nil {
		c.Logger.Error(err.Error())
	}
	c.Config = kafka.ConfigMap{}
	for key, val := range viper.GetStringMap("kafka-consumer") {
		c.Config[key] = val
	}
	c.Logger.Info(c.Config)
}

func (c *KafkaConsumer) StartConsumer() {
	c.ReadConfig()
	var err error
	configByte, _ := json.Marshal(c.Config)
	c.Logger.Info("Create consumer with config: ", string(configByte))
	c.Consumer, err = kafka.NewConsumer(&c.Config)
	if err != nil {
		c.Logger.Errorf("Create consumer failed: %v", err)
	}
	c.Logger.Info("Created consumer ", c.Consumer.String())
}

func (c *KafkaConsumer) rebalanceCallback(consumer *kafka.Consumer, event kafka.Event) error {
	switch event.(type) {
	case kafka.AssignedPartitions:
		assignedPartitions := event.(kafka.AssignedPartitions)
		c.Logger.Infof("Partitions were revoked: %v", assignedPartitions.Partitions)
		c.topicPartition = assignedPartitions.Partitions
	case kafka.RevokedPartitions:
		revokedPartitions := event.(kafka.RevokedPartitions)
		c.Logger.Infof("Partitions were revoked: %v", revokedPartitions.Partitions)
		topicPartition, err := consumer.Commit()
		if err != nil {
			c.Logger.Errorf("commit failed %v", err)
		}
		c.topicPartition = topicPartition
	}
	return nil
}

func (c *KafkaConsumer) SubscriptTopics() error {
	return c.Consumer.SubscribeTopics(c.Topics, c.rebalanceCallback)
}

type ProcessMessageFn func(message *kafka.Message) error

func (c *KafkaConsumer) processMessage(messageChans []<-chan *kafka.Message, processFn ProcessMessageFn) {
	for _, messageChan := range messageChans {
		go func(messageChan <-chan *kafka.Message) {
			for msg := range messageChan {
				retry := 0
				for retry < 10 {
					retry++
					if err := processFn(msg); err != nil {
						c.Logger.Errorf("process message error: %v, key: %s, value: %s, topicpartion: %v with error %v", msg, string(msg.Key), string(msg.Value), msg.TopicPartition, err)
						time.Sleep(time.Second)
					}
					break
				}
				go func() {
					_, err := c.Consumer.StoreMessage(msg)
					if err != nil {
						c.Logger.Errorf("commit failed due to %v, the next commit will serve as retry", err)
					}
				}()
			}
		}(messageChan)
	}
}

func (c *KafkaConsumer) GetCurrentPartition() (int, error) {
	numPartition := 0
	var consumerMetaData *kafka.Metadata
	var err error

	numRetry := 0
	for numRetry < 10 {
		numRetry++
		if numRetry >= 10 {
			return 0, err
		}
		consumerMetaData, err = c.Consumer.GetMetadata(nil, true, 10000)
		if err != nil {
			c.Logger.Errorf("Can't get consumer data due to error %v", err)
		}
		break
	}

	for _, topicMetaData := range consumerMetaData.Topics {
		numPartition += len(topicMetaData.Partitions)
	}
	return numPartition, nil
}

func (c *KafkaConsumer) ListenAndProcess() {
	defer c.Consumer.Close()
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	var numChan int
	numPartition, err := c.GetCurrentPartition()
	if err != nil {
		c.Logger.Infof("Due to the metadata couldn't be gotten from broker, the default max partition will be used")
		numChan = MaxPartition
	} else {
		numChan = numPartition
	}
	var messageChan []chan *kafka.Message
	for i := 0; i < numChan; i++ {
		messageChan[i] = make(chan *kafka.Message, 1000)
	}

	running := true
	for running {
		select {
		case sig := <-sigchan:
			c.Logger.Infof("Caught signal %v: terminating\n", sig)
			running = false
		default:
			message, err := c.Consumer.ReadMessage(5 * time.Second)
			if err != nil {
				c.Logger.Errorf("read message with error %v", err)
				continue
			}
			messageChan[message.TopicPartition.Partition] <- message
		}

	}
	for {
		if _, err := c.Consumer.Commit(); err != nil {
			c.Logger.Errorf("commit failed due to %v, will retry soon", err)
			time.Sleep(1 * time.Second)
		}
		break
	}
}
