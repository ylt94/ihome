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
		ctx.Set("user_info", rsp.User)
		ctx.Next()
	}
}
