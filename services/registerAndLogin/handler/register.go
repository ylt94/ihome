package handler

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/micro/go-micro/v2"
	"github.com/ylt94/ihome/services/registerAndLogin/model"
	"regexp"

	log "github.com/micro/go-micro/v2/logger"

	register "github.com/ylt94/ihome/services/registerAndLogin/proto/register"
	checkCaptcha "github.com/ylt94/ihome/services/getCaptcha/proto/checkCaptcha"
	checkSMS "github.com/ylt94/ihome/services/sendSMS/proto/checkSMS"
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
	captchaClient := checkCaptcha.NewCheckCaptchaService("go.micro.service.getCaptcha", microObj.Client())
	_, err := captchaClient.Check(ctx, &checkCaptcha.Request{Number: req.CaptchaCode, Uuid: req.Uuid})
	if err != nil {
		rsp.Status = register.RegisterStatus_Fail
		return err
	}

	//验证手机验证码是否正确
	checkSMSClient := checkSMS.NewCheckSMSService("go.micro.service.sendSMS", microObj.Client())
	_, err = checkSMSClient.Check(ctx, &checkSMS.Request{Phone: req.Phone, Code: req.SMSCode, Type: 0})
	if err != nil {
		return nil
	}

	//数据库操作
	//密码加密 MD5
	user := &model.User{Mobile: req.Phone, Name: req.Phone}
	M := md5.New()
	M.Write([]byte(req.Pwd))
	user.PasswordHash = hex.EncodeToString(M.Sum(nil))
	user.Register()

	return nil
}

