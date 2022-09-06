package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ylt94/ihome/web/controllers"
)

func orderRouter(R *gin.RouterGroup) {
	OR := R.Group("/order")
	{
		OR.POST("", controllers.OrderBuy)

		OR.PUT("/:order_id/status", controllers.OrderUpdate)


		OR.PUT("/:order_id/comment", controllers.OrderComment)
	}
}
