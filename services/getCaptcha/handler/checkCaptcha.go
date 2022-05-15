package handler

import (
	"context"
	"github.com/garyburd/redigo/redis"
	checkCaptcha "ihome/services/getCaptcha/proto/checkCaptcha"
	log "github.com/micro/go-micro/v2/logger"
	"ihome/services/getCaptcha/utils"
)
type CheckCaptcha struct {

}

func (e *CheckCaptcha) Check(Ctx context.Context, req *checkCaptcha.Request, rsp *checkCaptcha.Response) error {
	log.Info("Received CheckCaptcha.Call request")
	client, err := utils.Redis()
	if err != nil {
		return err
	}

	val, err := client.Do("get", req.Uuid)
	if err != nil {
		return err
	}
	code, err := redis.String(val, err)
	if err != nil {
		return err
	}

	if code == "" {
		rsp.Status = checkCaptcha.EheckStatus_Expire
	} else if code != req.Number {
		rsp.Status = checkCaptcha.EheckStatus_Fail
	} else {
		rsp.Status = checkCaptcha.EheckStatus_Succ
	}

	return nil
}
