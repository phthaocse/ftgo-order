package restauarant

import (
	"ftgo-order/internal/core/model"
	"ftgo-order/pkg/message"
)

type Created struct {
	name    string
	address string
	menu    []*model.MenuItem
}

func (e Created) GetEvent() string {
	return "create_restaurant"
}

func (e Created) GetMenu() []*model.MenuItem {
	return e.menu
}

func (e *Created) SetMenu(menu []*model.MenuItem) {
	e.menu = menu
}

func (e Created) GetAggregateId() string {
	return "restaurant"
}

func (e Created) GetMessage() message.Message {
	return message.Message{}
}
func (e Created) GetAggregateType() string {
	return "restaurant"
}

func (e Created) GetEventId() string {
	return "restaurant"
}
