package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"github.com/ylt94/ihome/services/index/handler"

	index "github.com/ylt94/ihome/services/index/proto/index"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.index"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	index.RegisterIndexHandler(service.Server(), new(handler.Index))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
