package kafka

import (
	"encoding/json"
	"ftgo-order/internal/outbound/interface/logger"
	"ftgo-order/pkg/helpers"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/viper"
)

type partitionInfo struct {
	topic         string
	partition     int
	currentOffset int
}

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

func (c *KafkaConsumer) ProcessMessage() {

}
