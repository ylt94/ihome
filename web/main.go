package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ylt94/ihome/web/routers"
)

func main() {
	R := gin.Default()
	R = routers.LoadRouters(R)
	R.Run()
}
