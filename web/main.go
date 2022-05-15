package main

import (
	"github.com/gin-gonic/gin"
	"ihome/web/routers"
)

func main() {
	R := gin.Default()
	R = routers.LoadRouters(R)
	R.Run()
}
