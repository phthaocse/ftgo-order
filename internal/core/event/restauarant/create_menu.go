package restauarant

import "ftgo-order/internal/core/model"

type Created struct {
	name    string
	address string
	menu    []*model.MenuItem
}

func (e Created) GetEvent() string {
	return e.name
}

func (e Created) GetMenu() []*model.MenuItem {
	return e.menu
}

func (e *Created) SetMenu(menu []*model.MenuItem) {
	e.menu = menu
}
