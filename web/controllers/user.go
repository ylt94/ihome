package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	auth "github.com/ylt94/ihome/services/user/proto/auth"
	user "github.com/ylt94/ihome/services/user/proto/user"
	"github.com/ylt94/ihome/web/config"
	"github.com/ylt94/ihome/web/utils"
)

func CheckRelaName(ctx *gin.Context) {
	userInfo := ctx.Get("user_info")

	client := micro.NewService()
	authService := auth.NewAuthService(config.USER, client)

	rsp, err := authService.CheckRealNameAuth(context.TODO(), &auth.CheckAuthRequest{UserId: userInfo.Id})
	if err != nil {
		msg := GetServiceError(err.Error())
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_UNKNOWERR, msg.Detail))
		return
	}

	//TODO 返回格式
}

func Auth(ctx *gin.Context) {
	userInfo := ctx.Get("user_info")
	client := micro.NewService()
	authService := auth.NewAuthService(config.USER, client)

	rsp, err := authService.RealNameAuth(context.TODO(), &auth.AuthRequest{})
	if err != nil {
		msg := GetServiceError(err.Error())
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_UNKNOWERR, msg.Detail))
		return
	}
}

func UserInfo(ctx *gin.Context) {
	client := micro.NewService()
	userService := user.NewUserService(config.USER, client)

}
