package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	house "github.com/ylt94/ihome/services/house/proto/house"
	"github.com/ylt94/ihome/web/config"
	"github.com/ylt94/ihome/web/utils"
)

type createParam struct {
	Title     string   `form:"title" binding:"required"`
	Price     uint32   `form:"price" binding:"required"`
	AreaId    uint32   `form:"area_id" binding:"required"`
	Address   string   `form:"address" binding:"required"`
	RoomCount uint32   `form:"room_count binding:"required""`
	Acreage   uint32   `form:"acreage" binding:"required"`
	Unit      string   `form:"uint" binding:"required"`
	Capacity  uint32   `form:"capacity" binding:"required"`
	Beds 	  string    `form:"beds" binding:"required"`
	Deposit   uint32	`form:"deposit binding:"required"`
	MinDays	  uint32    `form:"min_days" binding:"required"`
	MaxDays	  uint32    `form:"max_days" binding:"required"`
	Facility  []uint32 `form:"Facility[]""`
}

//房屋信息列表
func HouseList(ctx *gin.Context) {
	aid, _ := strconv.Atoi(ctx.Query("aid"))
	sd := ctx.Query("sd")
	ed := ctx.Query("ed")
	page, _ := strconv.Atoi(ctx.Query("p"))
	request := &house.ListRequest{Aid: uint32(aid), StartDate: sd, EndDate: ed, Page: uint32(page)}

	client := micro.NewService()
	houseService := house.NewHouseService(config.HOUSE, client.Client())

	rsp, err := houseService.List(context.TODO(), request)
	if err != nil {
		msg := GetServiceError(err.Error())
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_SERVERERR, "数据获取失败:"+msg.Detail))
		return
	}

	ctx.JSON(http.StatusOK, GetReturn(rsp, utils.RECODE_OK, "成功"))
	return
}

//上传房屋图片
func HouseUploadImage(ctx *gin.Context) {
	houseId := ctx.Param("house_id")
	houseIdInt, _ := strconv.ParseInt(houseId, 10, 32)
	if houseIdInt == 0 {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_PARAMERR, "参数错误:house_id"))
		return
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_PARAMERR, "图片保存失败:"+err.Error()))
		return
	}

	path := "/views/images/" + "house_" + houseId + "_" + file.Filename
	err = ctx.SaveUploadedFile(file, path)
	if err != nil {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_SERVERERR, "图片保存失败:"+err.Error()))
		return
	}

	client := micro.NewService()
	houseService := house.NewHouseService(config.HOUSE, client.Client())

	_, err = houseService.UploadImage(context.TODO(), &house.UploadImageRequest{HouseId: uint32(houseIdInt), Url: path})
	if err != nil {
		msg := GetServiceError(err.Error())
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_USERERR, "图片保存失败:"+msg.Detail))
		return
	}

	host := ctx.Request.Header.Get("Host")
	path = host + path

	res := struct{ Url string }{}
	res.Url = path
	ctx.JSON(http.StatusOK, GetReturn(res, utils.RECODE_OK, "成功"))
	return
}

//详情
func HouseDetail(ctx *gin.Context) {
	houseId := ctx.Param("house_id")
	houseIdInt, _ := strconv.Atoi(houseId)
	if houseIdInt == 0 {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_PARAMERR, "参数错误:house_id"))
		return
	}

	client := micro.NewService()
	houseService := house.NewHouseService(config.HOUSE, client.Client())

	rsp, err := houseService.Detail(context.TODO(), &house.DetailRequest{HouseId: uint32(houseIdInt)})
	if err != nil {
		msg := GetServiceError(err.Error())
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_SERVERERR, "数据获取失败:"+msg.Detail))
		return
	}
	res := struct {
		House  house.DetailResponse
		UserId uint32
	}{
		House:  *rsp,
		UserId: rsp.UserId,
	}
	ctx.JSON(http.StatusOK, GetReturn(res, utils.RECODE_OK, "成功"))
}

//发布房源
func HouseCreate(ctx *gin.Context) {
	var params createParam
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_PARAMERR, err.Error()))
		return
	}

	req := house.CreateRequest{
		Title: params.Title,
		Price: params.Price,
		AreaId: params.AreaId,
		Address: params.Address,
		RoomCount: params.RoomCount,
		Acreage: params.Acreage,
		Unit: params.Unit,
		Capacity: params.Capacity,
		Beds: params.Beds,
		Deposit: params.Deposit,
		MinDays: params.MinDays,
		MaxDays: params.MaxDays,
		Facility: params.Facility,
	}

	client := micro.NewService()
	houseService := house.NewHouseService(config.HOUSE, client.Client())

	rsp, err := houseService.Create(context.TODO(), &req)
	if err != nil {
		msg := GetServiceError(err.Error())
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_SERVERERR, msg.Detail))
		return
	}
	ctx.JSON(http.StatusOK, GetReturn(rsp, utils.RECODE_OK, "发布成功"))
	return
}
