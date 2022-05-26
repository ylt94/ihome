package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ylt94/ihome/web/controllers"
)

func registerAndLoginRouters(R *gin.RouterGroup) {
	R.GET("imagecode/:uuid", controllers.GetImageCode)
	R.GET("smscode/:phone", controllers.SendSMS)
	R.POST("users", controllers.Register)
	R.POST("sessions", controllers.Login)
}
