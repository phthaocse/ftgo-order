package event

import (
	"encoding/json"
	"ftgo-order/internal/outbound/interface/logger"
	"ftgo-order/pkg/message"
)

type DomainEventDispatcher struct {
	DispatcherId    string
	Handlers        []*DomainEventHandler
	MessageConsumer message.Consumer
	Logger          logger.Logger
}

func NewDomainEventDispatcher(messageConsumer message.Consumer, logger logger.Logger) *DomainEventDispatcher {
	return &DomainEventDispatcher{
		MessageConsumer: messageConsumer,
		Logger:          logger,
	}
}

func (d *DomainEventDispatcher) Dispatch() message.DispatcherFn {
	return func(message message.Message) {
		var header map[string][]byte
		if err := json.Unmarshal(message.Header, &header); err != nil {
			d.Logger.Errorf("can't read header: %v", err)
			return
		}
		d.Logger.Infof("Catch event %s", header["key"])
		return
	}
}

func (d *DomainEventDispatcher) Subscribe(subscriptionId string, channels map[string]struct{}, handler []*DomainEventHandler) {
	d.Handlers = handler
	d.MessageConsumer.Subscribe(subscriptionId, channels, d.Dispatch())
}
