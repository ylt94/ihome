package routers

import (
	"github.com/gin-gonic/gin"
)

func LoadRouters(R *gin.Engine) *gin.Engine{
	staticRoutes(R)
	api := R.Group("api/v1.0")
	{
		testRoutes(api)
		userRoutes(api)
		registerAndLoginRouters(api)
	}

	return R
}
