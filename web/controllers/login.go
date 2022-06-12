package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	login "github.com/ylt94/ihome/services/registerAndLogin/proto/login"
	"github.com/ylt94/ihome/web/config"
	"github.com/ylt94/ihome/web/utils"
	"net/http"
)

type loginParams struct {
	Phone string `json:"mobile" binding:"required"`
	Pwd string `json:"password" binding:"required"`
}

func Login(ctx *gin.Context) {
	var params loginParams
	ctx.ShouldBind(&params)

	microObj := micro.NewService()
	client := login.NewLoginService(config.REGISTER_AND_LOGIN, microObj.Client())
	resp, err := client.Login(context.TODO(), &login.LoginRequest{Pwd: params.Pwd, Phone: params.Phone})
	if err != nil {
		msg := GetServiceError(err.Error())
		fmt.Println(11111)
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_LOGINERR, msg.Detail))
		return
	}
	token := resp.Token

	//TODO session-cookie-token
	ctx.SetCookie("user_cookie", token, 0, "", "", true, true)
	ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_OK, "登录成功!"))
	return
}