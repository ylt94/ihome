package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	checkSMS "ihome/services/sendSMS/proto/checkSMS"
)

type CheckSMS struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *CheckSMS) Check(ctx context.Context, req *checkSMS.Request, rsp *checkSMS.Response) error {
	log.Info("Received SendSMS.Call request")
	return nil
}
