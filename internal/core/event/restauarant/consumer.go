package restauarant

import (
	"fmt"
	"ftgo-order/pkg/event"
)

type RestaurantEventConsumer struct {
}

func (c *RestaurantEventConsumer) DomainEventHandlers() []*event.DomainEventHandler {
	return []*event.DomainEventHandler{
		event.NewDomainEventHandler("restaurant", "create_restaurant", c.CreateRestaurant()),
	}
}

func (c *RestaurantEventConsumer) CreateRestaurant() event.HandlerFn {
	return func(event event.DomainEvent) {
		fmt.Println("confirmed")
	}
}
