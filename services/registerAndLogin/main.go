package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	handler "github.com/ylt94/ihome/services/registerAndLogin/handler"

	register "github.com/ylt94/ihome/services/registerAndLogin/proto/register"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.register"),
		micro.Version("latest"),
	)

	// Initialise service
	//service.Init()

	// Register Handler
	register.RegisterRegisterHandler(service.Server(), new(handler.Register))


	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
