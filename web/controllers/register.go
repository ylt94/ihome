package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	getCaptcha "github.com/services/getCaptcha/proto/getCaptcha"
	sendSMS "github.com/services/sendSMS/proto/sendSMS"
	"ihome/web/utils"
	"image/png"
	"net/http"
	"regexp"
)

//获取图片验证码
func GetImageCode(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	microObj := micro.NewService()

	client := getCaptcha.NewGetCaptchaService("go.micro.service.getCaptcha", microObj.Client())

	resp, err := client.Call(context.TODO(), &getCaptcha.Request{Number: "",CacheLiveTime: 5*60, Uuid: uuid})
	if err != nil {
		data := getReturn("", utils.RECODE_UNKNOWERR, "验证码生成失败")
		ctx.JSON(http.StatusOK, data)
		return
	}

	var img captcha.Image
	err = json.Unmarshal(resp.Image, &img)
	if err != nil {
		fmt.Println(err.Error())
		data := getReturn("", utils.RECODE_SERVERERR, "图片解析失败:"+err.Error())
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
		data := getReturn("", utils.RECODE_MOBILEERR, "请输入手机号")
		ctx.JSON(http.StatusOK, data)
		return
	}

	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$";
	reg := regexp.MustCompile(regular)
	if !reg.MatchString(phone) {
		ctx.JSON(http.StatusOK,getReturn("", utils.RECODE_MOBILEERR, "手机号格式错误!"))
		return
	}

	microObj := micro.NewService()

	client := sendSMS.NewSendSMSService("go.micro.service.sendSMS", microObj.Client())

	client.Send(context.TODO(), &sendSMS.Request{Phone: phone})
}