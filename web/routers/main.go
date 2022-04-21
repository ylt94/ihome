package routers

import (
	"github.com/gin-gonic/gin"
)

func LoadRouters(R *gin.Engine) *gin.Engine{
	testRoutes(R)
	return R
}
