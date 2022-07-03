package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	auth "github.com/ylt94/ihome/services/user/proto/auth"
	user "github.com/ylt94/ihome/services/user/proto/user"
	"github.com/ylt94/ihome/web/config"
	"github.com/ylt94/ihome/web/utils"
)

//检查用户实名信息
func UserCheckRelaName(ctx *gin.Context) {
	userInfo, exists := ctx.Get("user_info")
	if !exists {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_SESSIONERR, "请先登录"))
		return
	}

	client := micro.NewService()
	authService := auth.NewAuthService(config.USER, client.Client())

	rsp, err := authService.CheckRealNameAuth(context.TODO(), &auth.CheckAuthRequest{UserId: userInfo.Id})
	if err != nil {
		msg := GetServiceError(err.Error())
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_UNKNOWERR, msg.Detail))
		return
	}

	ctx.JSON(http.StatusOK, GetReturn(rsp, utils.RECODE_OK, "成功"))
	return
}

//实名认证
func UserAuth(ctx *gin.Context) {
	userInfo, exists := ctx.Get("user_info")
	if !exists {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_SESSIONERR, "请先登录"))
		return
	}

	realName := ctx.DefaultPostForm("real_name", "")
	idCard := ctx.DefaultPostForm("id_card", "")
	if realName == "" || idCard == "" {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_PARAMERR, "请将信息填写完整"))
		return
	}

	client := micro.NewService()
	authService := auth.NewAuthService(config.USER, client.Client())

	rsp, err := authService.RealNameAuth(context.TODO(), &auth.AuthRequest{UserId: userInfo.Id, RealName: realName, IdCard: idCard})
	if err != nil {
		msg := GetServiceError(err.Error())
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_UNKNOWERR,  "认证失败:"+msg.Detail))
		return
	}
	if rsp.Status == auth.AuthStatus_AuthFail {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_UNKNOWERR, "认证失败"))
		return
	}

	ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_OK, "成功"))
	return
}

//获取用户信息
func UserInfo(ctx *gin.Context) {
	userInfo, exists := ctx.Get("user_info")
	if !exists {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_SESSIONERR, "请先登录"))
		return
	}

	client := micro.NewService()
	userService := user.NewUserService(config.USER, client.Client())
	rsp, err := userService.Info(context.TODO(), &user.InfoRequest{userInfo.Id})
	if err != nil {
		msg := GetServiceError(err.Error())
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_UNKNOWERR, msg.Detail))
		return
	}

	ctx.JSON(http.StatusOK, GetReturn(rsp, utils.RECODE_OK, "成功"))
	return
}

//更新头像
func UploadUserAvatar(ctx *gin.Context) {
	userInfo ,exists := ctx.Get("user_info")
	if !exists {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_SESSIONERR, "请先登录"))
		return
	}

	userId := userInfo.Id

	//保存图片
	file, err := ctx.FormFile("avatar")
	if err != nil {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_PARAMERR, "头像上传失败:"+err.Error()))
		return
	}
	path := "/views/images/" +"user_avatar_"+strconv.Itoa(userId)+"_"+file.Filename
	err = ctx.SaveUploadedFile(file, "."+path)
	if err != nil {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_IOERR, "头像保存失败"))
		return
	}

	client := micro.NewService()
	userService := user.NewUserService(config.USER, client.Client())

	rsp, err := userService.UploadAvatar(context.TODO(), &user.AvatarRequest{UserId: userId, AvatarUrl: path})
	if err != nil {
		msg := GetServiceError(err.Error())
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_IOERR, "头像保存失败:"+msg.Detail))
		return
	}
	if rsp.Status == user.AvatarStatus_AvatarStatusFail {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_IOERR, "头像保存失败"))
		return
	}

	res := struct{AvatarUrl string}{}
	host := ctx.Request.Header.Get("Host")
	res.AvatarUrl = host+path
	ctx.JSON(http.StatusOK, GetReturn(res, utils.RECODE_OK, "成功"))
	return
}

//更新用户名
func UpdateUserName(ctx *gin.Context) {
	userInfo ,exists := ctx.Get("user_info")
	if !exists {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_SESSIONERR, "请先登录"))
		return
	}

	name := ctx.DefaultPostForm("name", "")
	if name == "" {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_PARAMERR, "请输入用户名"))
		return
	}

	client := micro.NewService()
	userService := user.NewUserService(config.USER, client.Client())
	rsp, err := userService.UpdateName(context.TODO(), &user.UpNameRequest{UserId: userInfo.Id, Name: name})
	if err != nil {
		msg := GetServiceError(err.Error())
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_IOERR, "更新失败:"+msg.Detail))
		return
	}

	if rsp.Status == user.UpNameStatus_UpNameFail {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_IOERR, "更新失败"))
		return
	}

	res := struct{Name string}{}
	res.Name = name
	ctx.JSON(http.StatusOK, GetReturn(res, utils.RECODE_OK, "成功"))
	return
}