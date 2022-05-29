package handler

import (
	"context"
	"errors"
	"github.com/micro/go-micro/v2"
	"github.com/ylt94/ihome/services/registerAndLogin/model"
	"github.com/ylt94/ihome/services/registerAndLogin/utils"
	"regexp"

	log "github.com/micro/go-micro/v2/logger"

	checkCaptcha "github.com/ylt94/ihome/services/getCaptcha/proto/checkCaptcha"
	register "github.com/ylt94/ihome/services/registerAndLogin/proto/register"
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

	captchaMicroObj := micro.NewService()
	//验证图片验证码是否正确
	captchaClient := checkCaptcha.NewCheckCaptchaService("go.micro.service.getCaptcha", captchaMicroObj.Client())
	_, err := captchaClient.Check(ctx, &checkCaptcha.Request{Number: req.CaptchaCode, Uuid: req.Uuid})
	if err != nil {
		rsp.Status = register.RegisterStatus_Fail
		return err
	}

	SMSMicroObj := micro.NewService()
	//验证手机验证码是否正确
	checkSMSClient := checkSMS.NewCheckSMSService("go.micro.service.sendSMS", SMSMicroObj.Client())
	_, err = checkSMSClient.Check(ctx, &checkSMS.Request{Phone: req.Phone, Code: req.SMSCode, Type: 0})
	if err != nil {
		return err
	}
	//数据库操作
	user := &model.User{}
	//检查手机号是否被注册
	err = user.GetUserByPhone(req.Phone)
	if err != nil {
		return nil
	}

	if user.Id != 0 {
		return errors.New("该手机号已被注册")
	}

	//密码加密 MD5
	user.Mobile = req.Phone
	user.Name = req.Phone
	user.PasswordHash = utils.EncryptionByMD5(req.Pwd)
	user.Register()
	return nil
}

