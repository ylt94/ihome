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
)

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

func SendSMS (ctx *gin.Context) {
	phone := ctx.Param("phone")
	if len(phone) == 0 {
		fmt.Println("请输入手机号")
		data := getReturn("", utils.RECODE_MOBILEERR, "请输入手机号")
		ctx.JSON(http.StatusOK, data)
		return
	}
	//TODO 正则验证手机号

	microObj := micro.NewService()

	client := sendSMS.NewSendSMSService("go.micro.service.sendSMS", microObj.Client())

	client.Send(context.TODO(), &sendSMS.Request{Phone: phone})
}