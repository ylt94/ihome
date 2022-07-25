package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ylt94/ihome/web/controllers"
)

func houseRouter(R *gin.RouterGroup) {
	HR := R.Group("/houses")
	{
		HR.POST("", controllers.HouseCreate)
		HR.POST("/:house_id/images", controllers.HouseUploadImage)
		HR.GET("/:house_id", controllers.HouseDetail)
		HR.GET("", controllers.HouseList)
	}
}
