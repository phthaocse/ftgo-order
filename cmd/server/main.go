package main

import (
	"ftgo-order/pkg/adapter/inbound/rest"
	postgresDb "ftgo-order/pkg/adapter/outbound/postgres_db"
	"ftgo-order/pkg/core/service"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	pgConn, err := postgresDb.Init()
	if err != nil {
		return
	}
	orderRepo := postgresDb.NewOrderRepo(pgConn)
	orderService := service.NewOrderService(orderRepo)
	services := service.BusinessService{
		OrderService: orderService,
	}
	ginServer := rest.NewGinServer()
	rest.StartHTTPServer(ginServer, services)
}
