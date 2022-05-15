package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	registerAndLogin "registerAndLogin/proto/registerAndLogin"
)

type RegisterAndLogin struct{}

func (e *RegisterAndLogin) Handle(ctx context.Context, msg *registerAndLogin.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *registerAndLogin.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
