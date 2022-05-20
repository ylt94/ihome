package handler

import (
	"context"
	"errors"
	"github.com/micro/go-micro/v2"
	"regexp"

	log "github.com/micro/go-micro/v2/logger"

	register "github.com/ylt94/ihome/services/registerAndLogin/proto/register"
	getCaptcha "github.com/ylt94/ihome/services/getCaptcha/proto/getCaptcha"
	sendSMS "github.com/ylt94/ihome/services/sendSMS/proto/checkSMS"
)

type Register struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Register) Register(ctx context.Context, req *register.Request, rsp *register.Response) error {
	log.Info("Received RegisterAndLogin.Call request")
	rsp.Status = register.RegisterStatus_Fail

	if req.ConfirmPwd != req.Pwd {
		return errors.New("两次输入的密码不一致!")
	}

	//验证手机格式
	regular :="^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	if !reg.MatchString(req.Phone) {
		return errors.New("手机号格式错误")
	}

	microObj := micro.NewService()

	//验证图片验证码是否正确
	client := getCaptcha.NewCheckCaptchaService("go.micro.service.getCaptcha", microObj.Client())
	_, err := client.Check(ctx, &checkCaptCha.Request{Number: req.CaptchaCode, Uuid: req.Uuid})
	if err != nil {
		rsp.Status = register.RegisterStatus_Fail
		return err
	}

	//验证手机验证码是否正确
	client := sendSMS.NewCheckSMSService("go.micro.service.sendSMS", microObj.Client())
	_, err = client.Check(ctx, &checkSMS.Request{Phone: req.Phone, Code: req.SMSCode, Type: 0})

	//数据库操作

	return nil
}

