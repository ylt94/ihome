package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	sendSMS "ihome/services/sendSMS/proto/sendSMS"
)

type SendSMS struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *SendSMS) Send(ctx context.Context, req *sendSMS.Request, rsp *sendSMS.Response) error {
	log.Info("Received SendSMS.Call request")
	return nil
}
