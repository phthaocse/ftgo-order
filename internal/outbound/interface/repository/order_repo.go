package repository

import (
	"ftgo-order/internal/core/model"
)

type OrderRepo interface {
	Create()
	GetById(id int)
	Update(order model.Order)
}
