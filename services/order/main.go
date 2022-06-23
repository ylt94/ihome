package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"github.com/ylt94/ihome/services/order/handler"

	order "github.com/ylt94/ihome/services/order/proto/order"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.order"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	order.RegisterOrderHandler(service.Server(), new(handler.Order))
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
