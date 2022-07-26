package postgres_db

import (
	"ftgo-order/pkg/core/model"
	"ftgo-order/pkg/interface/outbound/repository"
	"github.com/jackc/pgconn"
)

type orderRepo struct {
	pgConn *pgconn.PgConn
}

func NewOrderRepo(pgConn *pgconn.PgConn) repository.OrderRepo {
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
