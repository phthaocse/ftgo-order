package kafka

import (
	"encoding/json"
	"ftgo-order/internal/outbound/interface/logger"
	"ftgo-order/pkg/helpers"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/viper"
)

type KafkaProducer struct {
	Producer       *kafka.Producer
	topics         []string
	Config         kafka.ConfigMap
	Logger         logger.Logger
	topicPartition []kafka.TopicPartition
}

func (c *KafkaProducer) ReadConfig() {
	viper.SetConfigName("kafka")
	viper.SetConfigType("yaml")
	projectPath := helpers.ProjectPath()
	viper.AddConfigPath(projectPath + "config")
	if err := viper.ReadInConfig(); err != nil {
		c.Logger.Error(err.Error())
	}
	c.Config = kafka.ConfigMap{}
	for key, val := range viper.GetStringMap("kafka-event_publisher") {
		c.Config[key] = val
	}
	c.Logger.Info(c.Config)
}

func NewProducer(logger logger.Logger) *KafkaProducer {
	producer := KafkaProducer{}
	producer.Logger = logger
	producer.ReadConfig()
	var err error
	configByte, _ := json.Marshal(producer.Config)
	producer.Logger.Info("Create event_publisher with config: ", string(configByte))
	producer.Producer, err = kafka.NewProducer(&producer.Config)
	if err != nil {
		producer.Logger.Errorf("Create event_publisher failed: %v", err)
	}
	producer.Logger.Info("Created event_publisher ", producer.Producer.String())
	return &producer
}
