package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	handler "github.com/ylt94/ihome/services/registerAndLogin/handler"

	register "github.com/ylt94/ihome/services/registerAndLogin/proto/register"
	login "github.com/ylt94/ihome/services/registerAndLogin/proto/login"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.registerAndLogin"),
		micro.Version("latest"),
	)

	// Initialise service
	//service.Init()

	// Register Handler
	register.RegisterRegisterHandler(service.Server(), new(handler.Register))
	// Login Handler
	login.RegisterLoginHandler(service.Server(), new(handler.Login))
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
