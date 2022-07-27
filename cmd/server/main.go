package main

import (
	"ftgo-order/pkg/core/service"
	rest2 "ftgo-order/pkg/inbound/adapter/rest"
	coreRepo "ftgo-order/pkg/outbound/adapter/core_repo"
	"ftgo-order/pkg/outbound/adapter/logger"
	"ftgo-order/pkg/outbound/adapter/postgres_repo"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	pgConn, err := postgres_repo.Init(logger.ZapLogger)
	if err != nil {
		return
	}
	orderPostgresRepo := postgres_repo.NewOrderPostgresRepo(pgConn)
	orderRepo := coreRepo.NewOrderRepo(orderPostgresRepo)
	orderService := service.NewOrderService(orderRepo)
	services := service.BusinessService{
		OrderService: orderService,
	}
	ginServer := rest2.NewGinServer(logger.ZapLogger)
	rest2.StartHTTPServer(ginServer, services)
}
