package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	index "github.com/ylt94/ihome/services/index/proto/index"
	"github.com/ylt94/ihome/web/config"
	"github.com/ylt94/ihome/web/utils"
)

func Areas(ctx *gin.Context) {

	client := micro.NewService()
	indexService := index.NewIndexService(config.INDEX, client)

	rsp, err := indexService.Area(context.TODO(), &index.AreaRequest{})
	if err != nil {
		msg := GetServiceError(err.Error())
		data := GetReturn("", utils.RECODE_UNKNOWERR, "获取地区信息失败:"+msg)
		ctx.JSON(http.StatusOK, data)
		return
	}

	data := ctx.JSON(http.StatusOK, GetReturn(rsp.Areas), utils.RECODE_OK, "成功")
	return
}

func Banner(ctx *gin.Context) {
	client := micro.NewService()
	indexService := index.NewIndexService(config.INDEX, client)

	rsp, err := indexService.Banner(context.TODO(), &index.BannerRequest{})
	if err != nil {
		msg := GetServiceError(err.Error())
		data := GetReturn("", utils.RECODE_UNKNOWERR, "获取首页Banner失败:"+msg)
		ctx.JSON(http.StatusOK, data)
		return
	}

	data := ctx.JSON(http.StatusOK, GetReturn(rsp.Houses), utils.RECODE_OK, "成功")
	return
}
