package restauarant

import (
	"fmt"
	"ftgo-order/pkg/event"
)

type RestaurantEventConsumer struct {
}

func (c *RestaurantEventConsumer) DomainEventHandlers() []*event.DomainEventHandler {
	createdEvent := Created{}
	return []*event.DomainEventHandler{
		event.NewDomainEventHandler("restaurant", createdEvent, c.CreateRestaurant()),
	}
}

func (c *RestaurantEventConsumer) CreateRestaurant() event.HandlerFn {
	return func(event event.DomainEvent) {
		fmt.Println("confirmed")
	}
}
