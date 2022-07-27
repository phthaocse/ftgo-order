package main

import (
	service2 "ftgo-order/internal/core/service"
	"ftgo-order/internal/inbound/adapter/consumer/kafka"
	"ftgo-order/internal/inbound/adapter/rest"
	coreRepo "ftgo-order/internal/outbound/adapter/core_repo"
	"ftgo-order/internal/outbound/adapter/logger"
	postgres_repo2 "ftgo-order/internal/outbound/adapter/postgres_repo"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	kafka.StartConsumer(logger.ZapLogger)
	pgConn, err := postgres_repo2.Init(logger.ZapLogger)
	if err != nil {
		return
	}
	orderPostgresRepo := postgres_repo2.NewOrderPostgresRepo(pgConn)
	orderRepo := coreRepo.NewOrderRepo(orderPostgresRepo)
	orderService := service2.NewOrderService(orderRepo)
	services := service2.BusinessService{
		OrderService: orderService,
	}
	ginServer := rest.NewGinServer(logger.ZapLogger)
	rest.StartHTTPServer(ginServer, services)
}
