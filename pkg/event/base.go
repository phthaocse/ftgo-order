package event

import "ftgo-order/pkg/message"

type DomainEvent interface {
	GetEvent() string
	GetAggregateId() string
	GetMessage() message.Message
	GetAggregateType() string
	GetEventId() string
}
