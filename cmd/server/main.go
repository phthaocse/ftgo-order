package main

import (
	service2 "ftgo-order/internal/core/service"
	"ftgo-order/internal/inbound/adapter/consumer/kafka"
	"ftgo-order/internal/inbound/adapter/rest"
	coreRepo "ftgo-order/internal/outbound/adapter/core_repo"
	"ftgo-order/internal/outbound/adapter/logger"
	postgres_repo2 "ftgo-order/internal/outbound/adapter/postgres_repo"
	"github.com/spf13/viper"
	"os"
)

func main() {
	viper.AutomaticEnv()
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
	if len(os.Args) > 1 && os.Args[1] == "order-consumer" {
		orderConsumer := kafka.NewOrderConsumer(logger.ZapLogger)
		orderConsumer.StartOrderConsumer()
	} else {
		ginServer := rest.NewGinServer(logger.ZapLogger)
		rest.StartHTTPServer(ginServer, services)
	}

}
