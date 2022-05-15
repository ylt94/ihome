package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	sendSMS "sendSMS/proto/sendSMS"
)

type SendSMS struct{}

func (e *SendSMS) Handle(ctx context.Context, msg *sendSMS.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *sendSMS.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
