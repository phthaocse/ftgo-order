package event

import "ftgo-order/pkg/message"

type DomainEvent interface {
	GetEvent() string
	GetAggregateId() string
	GetMessage() message.Message
	GetAggregateType() string
	GetEventId() string
}

type Handler interface {
	ServeEvent(event DomainEvent)
}

type HandlerFn func(event DomainEvent)

func (f HandlerFn) ServeEvent(event DomainEvent) {
	f(event)
}

type DomainEventHandler struct {
	aggregateType string
	event         string
	handler       Handler
}

func NewDomainEventHandler(aggregateType string, event string, handler Handler) *DomainEventHandler {
	return &DomainEventHandler{
		aggregateType: aggregateType,
		event:         event,
		handler:       handler,
	}
}
