package postgres_repo

import (
	"ftgo-order/internal/core/model"
	"github.com/jackc/pgconn"
)

type OrderPostgresRepo struct {
	pgConn *pgconn.PgConn
}

func NewOrderPostgresRepo(pgConn *pgconn.PgConn) *OrderPostgresRepo {
	return &OrderPostgresRepo{
		pgConn: pgConn,
	}
}

func (r *OrderPostgresRepo) Create() {

}

func (r *OrderPostgresRepo) GetById(id int) {

}

func (r *OrderPostgresRepo) Update(order model.Order) {

}
