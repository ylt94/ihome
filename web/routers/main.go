package routers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	login "github.com/ylt94/ihome/services/registerAndLogin/proto/login"
	"github.com/ylt94/ihome/web/config"
	"github.com/ylt94/ihome/web/controllers"
	"github.com/ylt94/ihome/web/utils"
	"net/http"
)

type User struct{
	Id int `json:"user_id"`
	Name string `json:"name"`
	Mobile string `json:"mobile"`
	Avatar_url string `json:"avatar_url"`
}

func LoadRouters(R *gin.Engine) *gin.Engine{
	staticRoutes(R)
	api := R.Group("api/v1.0")
	{
		testRoutes(api)
		registerAndLoginRouters(api)
		api.Use(AuthUserInfo())
		{
			userRoutes(api)
		}
	}

	return R
}

func AuthUserInfo() gin.HandlerFunc{
	return func (ctx *gin.Context) {
		token, _ := ctx.Cookie("user_cookie")
		if token == "" {
			ctx.JSON(http.StatusOK, controllers.GetReturn("", utils.RECODE_SESSIONERR, "登录过期，请重新登录!"))
			return
		}

		microObj := micro.NewService()
		client := login.NewLoginService(config.REGISTER_AND_LOGIN, microObj.Client())

		rsp, err := client.Auth(context.TODO(), &login.AuthRequest{Token: token})
		if err != nil {
			msg := controllers.GetServiceError(err.Error())
			ctx.JSON(http.StatusOK, controllers.GetReturn("", utils.RECODE_SESSIONERR, msg.Detail))
			return
		}
		user := rsp.GetUser()
		userInfo := User{Id:int(user.Id), Name: user.Name, Mobile: user.Mobile, Avatar_url: user.AvatarUrl}
		ctx.Set("user_info", userInfo)
		ctx.Next()
	}
}
