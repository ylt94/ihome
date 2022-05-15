package routers

import (
	"github.com/gin-gonic/gin"
	"ihome/web/controllers"
)

func registerAndLoginRouters(R *gin.RouterGroup) {
	R.GET("imagecode/:uuid", controllers.GetImageCode)
	R.GET("smscode/:phone", controllers.SendSMS)
}
