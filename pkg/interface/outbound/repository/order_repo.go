package repository

import (
	"ftgo-order/pkg/core/model"
)

type OrderRepo interface {
	Create()
	GetById(id int)
	Update(order model.Order)
}
