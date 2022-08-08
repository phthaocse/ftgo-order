package producer

import "ftgo-order/internal/core/event"

type EventPublisher interface {
	Publish(aggregateType string, id interface{}, events []event.DomainEvent)
}
