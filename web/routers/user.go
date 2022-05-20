package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ylt94/ihome/web/utils"
	"net/http"
)

func userRoutes(R *gin.RouterGroup) {
	R.GET("token",func(ctx *gin.Context){
		res := map[string]string {
			"errno": utils.RECODE_SESSIONERR,
			"data": "",
		}
		ctx.JSON(http.StatusOK, res)
	})
}