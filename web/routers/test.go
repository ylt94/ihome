package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func testRoutes(R *gin.RouterGroup) {
	R.GET("/", func(c *gin.Context){
		c.String(http.StatusOK, "test")
	})
}