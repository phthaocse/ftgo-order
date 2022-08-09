package event

import (
	"ftgo-order/internal/core/event/restauarant"
	"ftgo-order/internal/core/service"
	"ftgo-order/pkg/event"
)

type orderEventConsumer struct {
	OrderService service.OrderServiceI
}

func NewOrderEventConsumer(orderService service.OrderServiceI) *orderEventConsumer {
	return &orderEventConsumer{OrderService: orderService}
}

func (e *orderEventConsumer) Handlers() []*event.DomainEventHandler {
	createMenuEvent := restauarant.Created{}
	return []*event.DomainEventHandler{
		event.NewDomainEventHandler("Restaurant", createMenuEvent, e.CreateMenu()),
	}
}

func (e *orderEventConsumer) CreateMenu() event.HandlerFn {
	return func(event event.DomainEvent) {
		event.GetAggregateType()
		e.OrderService.CreateMenu()
	}
}
