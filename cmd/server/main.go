package main

import (
	"ftgo-order/pkg/adapter/inbound/rest"
	coreRepo "ftgo-order/pkg/adapter/outbound/core_repo"
	"ftgo-order/pkg/adapter/outbound/logger"
	postgresRepo "ftgo-order/pkg/adapter/outbound/postgres_repo"
	"ftgo-order/pkg/core/service"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	pgConn, err := postgresRepo.Init(logger.ZapLogger)
	if err != nil {
		return
	}
	orderPostgresRepo := postgresRepo.NewOrderPostgresRepo(pgConn)
	orderRepo := coreRepo.NewOrderRepo(orderPostgresRepo)
	orderService := service.NewOrderService(orderRepo)
	services := service.BusinessService{
		OrderService: orderService,
	}
	ginServer := rest.NewGinServer()
	rest.StartHTTPServer(ginServer, services)
}
