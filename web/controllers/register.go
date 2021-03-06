package controllers

import (
	register "github.com/ylt94/ihome/services/registerAndLogin/proto/register"
	"context"
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	getCaptcha "github.com/ylt94/ihome/services/getCaptcha/proto/getCaptcha"
	sendSMS "github.com/ylt94/ihome/services/sendSMS/proto/sendSMS"
	"github.com/ylt94/ihome/web/config"
	"github.com/ylt94/ihome/web/utils"
	"image/png"
	"net/http"
	"regexp"

)

type registerParams struct {
	ConfirmPassword string `json:"confirm_password" binding:"required"`
	Phone string `json:"mobile" binding:"required"`
	SMSCode string `json:"sms_code" binding:"required"`
	Pwd string `json:"password" binding:"required"`
	Uuid string `json:"uuid" binding:"required"`
	ChaptchaCode string `json:"chaptcha_code" binding:"required"`
}

//获取图片验证码
func GetImageCode(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	microObj := micro.NewService()

	client := getCaptcha.NewGetCaptchaService(config.GET_CAPTCHA, microObj.Client())

	resp, err := client.Call(context.TODO(), &getCaptcha.Request{Number: "",CacheLiveTime: 5*60, Uuid: uuid})
	if err != nil {
		msg := GetServiceError(err.Error())
		data := GetReturn("", utils.RECODE_UNKNOWERR, "验证码生成失败:"+msg.Detail)
		ctx.JSON(http.StatusOK, data)
		return
	}

	var img captcha.Image
	err = json.Unmarshal(resp.Image, &img)
	if err != nil {
		msg := GetServiceError(err.Error())
		data := GetReturn("", utils.RECODE_SERVERERR, "图片解析失败:"+msg.Detail)
		ctx.JSON(http.StatusOK, data)
		return
	}

	png.Encode(ctx.Writer, img)
	return
}

//发送短信验证码
func SendSMS (ctx *gin.Context) {
	phone := ctx.Param("phone")
	if len(phone) == 0 {
		fmt.Println("请输入手机号")
		data := GetReturn("", utils.RECODE_MOBILEERR, "请输入手机号")
		ctx.JSON(http.StatusOK, data)
		return
	}

	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$";
	reg := regexp.MustCompile(regular)
	if !reg.MatchString(phone) {
		ctx.JSON(http.StatusOK,GetReturn("", utils.RECODE_MOBILEERR, "手机号格式错误!"))
		return
	}

	microObj := micro.NewService()

	client := sendSMS.NewSendSMSService(config.SEND_SMS, microObj.Client())
	_, err := client.Send(context.TODO(), &sendSMS.Request{Phone: phone})
	if err != nil {
		msg := GetServiceError(err.Error())
		ctx.JSON(http.StatusOK,GetReturn("", utils.RECODE_DATAERR, msg.Detail))
		return
	}

	ctx.JSON(http.StatusOK, GetReturn(""))
	return
}

//注册
func Register(ctx *gin.Context) {
	var params registerParams

	//TODO 自定义验证规则--正则验证手机号
	ctx.ShouldBind(&params)

	req := register.RegisterRequest{
		Phone: params.Phone,
		Pwd: params.Pwd,
		ConfirmPwd: params.ConfirmPassword,
		Uuid: params.Uuid,
		SMSCode: params.SMSCode,
		CaptchaCode: params.ChaptchaCode,
	}

	microObj := micro.NewService()
	client := register.NewRegisterService(config.REGISTER_AND_LOGIN, microObj.Client())
	_, err := client.Register(context.TODO(), &req)
	if err != nil {
		msg := GetServiceError(err.Error())
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_REGIERR, msg.Detail))
		return
	}
	ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_SESSIONERR))
	return
}