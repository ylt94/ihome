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

	UR := R.Group("/user")
	{
		UR.GET("", controllers.UserInfo)
		//上传头像
		UR.POST("/avatar", controllers.UploadUserAvatar)
		//修改用户名
		UR.PUT("/name", controllers.UpdateUserName)

		//获取（检查）用户实名信息
		UR.GET("/auth", controllers.UserCheckRelaName)

		//实名认证
		UR.POST("/auth", controllers.UserAuth)

		UR.GET("/house", controllers.HouseList)

		UR.GET("/orders", controllers.OrderList)
	}

}