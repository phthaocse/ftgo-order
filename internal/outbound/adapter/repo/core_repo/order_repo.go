package core_repo

import (
	"ftgo-order/internal/core/model"
	"ftgo-order/internal/outbound/adapter/repo/postgres_repo"
	"ftgo-order/internal/outbound/interface/repository"
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
