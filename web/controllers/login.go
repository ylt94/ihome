package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	login "github.com/ylt94/ihome/services/registerAndLogin/proto/login"
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
	client := login.NewLoginService("go.micro.service.registerAndLogin", microObj.Client())
	resp, err := client.Login(context.TODO(), &login.Request{Pwd: params.Pwd, Phone: params.Phone})
	if err != nil {
		ctx.JSON(http.StatusOK, getReturn("", utils.RECODE_LOGINERR, err.Error()))
		return
	}
	token := resp.Token
	fmt.Println(token)
	//TODO session-cookie-token
}