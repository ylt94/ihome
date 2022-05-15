package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"registerAndLogin/handler"
	"registerAndLogin/subscriber"

	registerAndLogin "registerAndLogin/proto/registerAndLogin"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.registerAndLogin"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	registerAndLogin.RegisterRegisterAndLoginHandler(service.Server(), new(handler.RegisterAndLogin))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.registerAndLogin", service.Server(), new(subscriber.RegisterAndLogin))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
