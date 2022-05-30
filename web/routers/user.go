package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ylt94/ihome/web/controllers"
	"github.com/ylt94/ihome/web/utils"
	"net/http"
)

func userRoutes(R *gin.RouterGroup) {
	R.GET("session",func(ctx *gin.Context){
		res, exists := ctx.Get("user_info")
		if !exists {
			ctx.JSON(http.StatusOK, controllers.GetReturn("", utils.RECODE_SESSIONERR, "请先登录"))
			return
		}
		ctx.JSON(http.StatusOK, controllers.GetReturn(res))
		return
	})
}