package kafka

import (
	"ftgo-order/internal/outbound/interface/logger"
	"ftgo-order/pkg/message"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type RestaurantMessageConsumer struct {
	Consumer   *KafkaConsumer
	Logger     logger.Logger
	Dispatcher message.Dispatcher
}

func NewRestaurantConsumer(logger logger.Logger) *RestaurantMessageConsumer {
	return &RestaurantMessageConsumer{
		Consumer: NewConsumer(logger),
		Logger:   logger,
	}
}

func (oc *RestaurantMessageConsumer) ProcessMessage(msg *kafka.Message) error {
	dispatchMsg := toMessage(msg)
	oc.Dispatcher.Dispatch(dispatchMsg)
	return nil
}

func (oc *RestaurantMessageConsumer) Subscribe(subscriberId string, channels map[string]struct{}, dispatcher message.Dispatcher) {
	err := oc.Consumer.SubscriptTopics(mapToSlice(channels))
	oc.Dispatcher = dispatcher
	if err != nil {
		oc.Logger.Errorf("can't subscript topic order %v", err)
		panic(err)
	}
}

func (oc *RestaurantMessageConsumer) Start() {
	oc.Consumer.ListenAndProcess(oc.ProcessMessage)
}
