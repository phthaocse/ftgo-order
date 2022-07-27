package kafka

import (
	"encoding/json"
	"ftgo-order/internal/outbound/interface/logger"
	"ftgo-order/pkg/helpers"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/viper"
)

type KafkaConsumer struct {
	Consumer *kafka.Consumer
	Config   kafka.ConfigMap
}

func (k *KafkaConsumer) ReadConfig(logger logger.Logger) {
	viper.SetConfigName("kafka")
	viper.SetConfigType("yaml")
	projectPath := helpers.ProjectPath()
	viper.AddConfigPath(projectPath + "config")
	if err := viper.ReadInConfig(); err != nil {
		logger.Error(err.Error())
	}
	k.Config = kafka.ConfigMap{}
	for key, val := range viper.GetStringMap("kafka-consumer") {
		k.Config[key] = val
	}
	logger.Info(k.Config)
}

func StartConsumer(logger logger.Logger) {
	consumer := KafkaConsumer{}
	consumer.ReadConfig(logger)
	var err error
	configByte, _ := json.Marshal(consumer.Config)
	logger.Info("Create consumer with config: ", string(configByte))
	consumer.Consumer, err = kafka.NewConsumer(&consumer.Config)
	if err != nil {
		logger.Errorf("Create consumer failed: %v", err)
	}
	logger.Info("Created consumer ", consumer.Consumer.String())
}
