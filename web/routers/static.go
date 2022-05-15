package routers

import "github.com/gin-gonic/gin"

func staticRoutes(R *gin.Engine) {
	R.Static("/view", "./views")
}
