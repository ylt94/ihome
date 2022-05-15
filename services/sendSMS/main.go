package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"ihome/services/sendSMS/handler"

	sendSMS "ihome/services/sendSMS/proto/sendSMS"
	checkSMS "ihome/services/sendSMS/proto/checkSMS"
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
	sendSMS.RegisterSendSMSHandler(service.Server(), new(handler.SendSMS))
	checkSMS.RegisterCheckSMSHandler(service.Server(), new(handler.CheckSMS))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
