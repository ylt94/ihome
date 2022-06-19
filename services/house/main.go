package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"github.com/ylt94/ihome/services/house/handler"
	house "github.com/ylt94/ihome/services/house/proto/house"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.house"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	house.RegisterHouseHandler(service.Server(), new(handler.House))
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
