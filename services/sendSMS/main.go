package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	handler "github.com/ylt94/ihome/services/sendSMS/handler"
	checkSMS "github.com/ylt94/ihome/services/sendSMS/proto/checkSMS"

	sendSMS "github.com/ylt94/ihome/services/sendSMS/proto/sendSMS"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.sendSMS"),
		micro.Version("latest"),
	)

	// Initialise service
	//service.Init()

	// Register Handler
	checkSMS.RegisterCheckSMSHandler(service.Server(), new(handler.CheckSMS))
	sendSMS.RegisterSendSMSHandler(service.Server(), new(handler.SendSMS))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
