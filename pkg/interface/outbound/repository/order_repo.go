package repository

import (
	"ftgo-order/pkg/core/model"
	"github.com/jackc/pgconn"
)

type OrderRepo interface {
	Create()
	GetById(id int)
	Update(order model.Order)
}

type orderRepo struct {
	pgConn *pgconn.PgConn
}

func NewOrderRepo(pgConn *pgconn.PgConn) OrderRepo {
	return &orderRepo{
		pgConn: pgConn,
	}
}

func (r *orderRepo) Create() {

}

func (r *orderRepo) GetById(id int) {

}

func (r *orderRepo) Update(order model.Order) {

}
