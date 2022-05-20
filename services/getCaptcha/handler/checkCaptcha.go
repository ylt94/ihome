package handler

import (
	"context"
	"errors"
	"github.com/garyburd/redigo/redis"
	log "github.com/micro/go-micro/v2/logger"
	checkCaptcha "github.com/ylt94/ihome/services/getCaptcha/proto/checkCaptcha"
	"github.com/ylt94/ihome/services/getCaptcha/utils"
)
type CheckCaptcha struct {

}

func (e *CheckCaptcha) Check(Ctx context.Context, req *checkCaptcha.Request, rsp *checkCaptcha.Response) error {
	log.Info("Received CheckCaptcha.Call request")
	client, err := utils.Redis()
	if err != nil {
		return err
	}
	defer client.Close()

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
		return errors.New("图形验证码已过期，请点击刷新")
	} else if code != req.Number {
		rsp.Status = checkCaptcha.EheckStatus_Fail
		return errors.New("图形验证码错误!")
	} else {
		rsp.Status = checkCaptcha.EheckStatus_Succ
		client.Do("del", req.Uuid)
		return nil
	}

	return nil
}
