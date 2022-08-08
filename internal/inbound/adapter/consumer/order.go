package consumer

import (
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
	return []*event.DomainEventHandler{
		event.NewDomainEventHandler("Restaurant", "create_menu", e.CreateMenu()),
	}
}

func (e *orderEventConsumer) CreateMenu() event.HandlerFn {
	return func(event event.DomainEvent) {
		event.GetAggregateType()
		e.OrderService.CreateMenu()
	}
}
