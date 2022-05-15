package handler

import (
	"context"
	checkCaptcha "ihome/services/getCaptcha/proto/checkCaptcha"
	log "github.com/micro/go-micro/v2/logger"
)
type CheckCaptcha struct {

}

func (e *CheckCaptcha) Check(Ctx context.Context, req *checkCaptcha.Request, rsp *checkCaptcha.Response) error {
	log.Info("Received CheckCaptcha.Call request")
	return nil
}
