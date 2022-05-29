package handler

import (
	"context"
	"github.com/ylt94/ihome/services/sendSMS/utils"

	log "github.com/micro/go-micro/v2/logger"

	checkSMS "github.com/ylt94/ihome/services/sendSMS/proto/checkSMS"
)

type CheckSMS struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *CheckSMS) Check(ctx context.Context, req *checkSMS.Request, rsp *checkSMS.Response) error {
	log.Info("Received CheckSMS.Call request")
	client := utils.Redis()
	defer client.Close()

	key := checkSMS.SMSType_name[int32(checkSMS.SMSType_Register)] + "_" + req.Phone
	_ ,err := client.Do("get", key)
	if err != nil {
		return err
	}
	//if val == "" {
	//	rsp.Status = checkSMS.EheckStatus_Expire
	//	return errors.New("手机验证码已过期!")
	//} else if val != string(req.Code) {
	//	rsp.Status = checkSMS.EheckStatus_Fail
	//	return errors.New("手机验证码已错误!")
	//} else if val == string(req.Code) {
		rsp.Status = checkSMS.EheckStatus_Succ
		client.Do("del", key)
		return nil
	//}

	return nil
}
