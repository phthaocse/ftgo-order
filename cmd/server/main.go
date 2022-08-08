package main

import (
	service2 "ftgo-order/internal/core/service"
	"ftgo-order/internal/inbound/adapter/consumer"
	"ftgo-order/internal/inbound/adapter/consumer/kafka"
	"ftgo-order/internal/inbound/adapter/rest"
	"ftgo-order/internal/outbound/adapter/logger"
	coreRepo "ftgo-order/internal/outbound/adapter/repo/core_repo"
	"ftgo-order/internal/outbound/adapter/repo/postgres_repo"
	"ftgo-order/pkg/event"
	"github.com/spf13/viper"
	"os"
)

func main() {
	viper.AutomaticEnv()
	pgConn, err := postgres_repo.Init(logger.ZapLogger)
	if err != nil {
		return
	}
	orderPostgresRepo := postgres_repo.NewOrderPostgresRepo(pgConn)
	orderRepo := coreRepo.NewOrderRepo(orderPostgresRepo)
	orderService := service2.NewOrderService(orderRepo)
	services := service2.BusinessService{
		OrderService: orderService,
	}

	if len(os.Args) > 1 && os.Args[1] == "order-consumer" {
		orderEventConsumer := consumer.NewOrderEventConsumer(orderService)
		orderEventConsumer.Handlers()
		restaurantMessageConsumer := kafka.NewRestaurantConsumer(logger.ZapLogger)
		restaurantDispatcher := event.NewDomainEventDispatcher(restaurantMessageConsumer, logger.ZapLogger)
		restaurantDispatcher.Subscribe("RES1", map[string]struct{}{"restaurant": {}}, orderEventConsumer.Handlers())
		restaurantMessageConsumer.Start()
	} else {
		ginServer := rest.NewGinServer(logger.ZapLogger)
		rest.StartHTTPServer(ginServer, services)
	}

}
