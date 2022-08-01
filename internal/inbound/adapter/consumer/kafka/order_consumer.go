package kafka

import (
	"ftgo-order/internal/outbound/interface/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type OrderConsumer struct {
	Consumer *KafkaConsumer
	Logger   logger.Logger
}

func NewOrderConsumer(logger logger.Logger) *OrderConsumer {
	return &OrderConsumer{
		Logger: logger,
	}
}

func (oc *OrderConsumer) ProcessMessage(msg *kafka.Message) error {
	oc.Logger.Info(msg)
	return nil
}

func (oc *OrderConsumer) StartOrderConsumer() {
	oc.Consumer = NewConsumer(oc.Logger)
	err := oc.Consumer.SubscriptTopic("test-topic")
	if err != nil {
		oc.Logger.Errorf("can't subscript topic order %v", err)
		panic(err)
	}
	oc.Consumer.ListenAndProcess(oc.ProcessMessage)
}
