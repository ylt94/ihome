package main

import (
	"ihome/services/getCaptcha/handler"
	getCaptcha "ihome/services/getCaptcha/proto/getCaptcha"
	checkCaptcha "ihome/services/getCaptcha/proto/checkCaptcha"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.getCaptcha"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	//生成验证码
	getCaptcha.RegisterGetCaptchaHandler(service.Server(), new(handler.GetCaptcha))
	//检查验证码
	checkCaptcha.RegisterCheckCaptchaHandler(service.Server(), new(handler.CheckCaptcha))
	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.service.getCaptcha", service.Server(), new(subscriber.GetCaptcha))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
