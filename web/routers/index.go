package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ylt94/ihome/web/controllers"
)

func indexRouter(R *gin.RouterGroup) {
	R.GET("/house/index", controllers.IndexBanner)
	R.GET("/areas", controllers.IndexAreas)
}