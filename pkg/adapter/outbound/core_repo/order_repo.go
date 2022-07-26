package core_repo

import (
	"ftgo-order/pkg/adapter/outbound/postgres_repo"
	"ftgo-order/pkg/core/model"
	"ftgo-order/pkg/interface/outbound/repository"
)

type orderRepo struct {
	orderPostgresRepo *postgres_repo.OrderPostgresRepo
}

func NewOrderRepo(orderPostgresRepo *postgres_repo.OrderPostgresRepo) repository.OrderRepo {
	return &orderRepo{orderPostgresRepo: orderPostgresRepo}
}

func (r *orderRepo) Create() {

}

func (r *orderRepo) GetById(id int) {

}

func (r *orderRepo) Update(order model.Order) {

}
