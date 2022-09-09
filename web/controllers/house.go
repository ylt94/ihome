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
	Title     string   `json:"title" form:"title" binding:"required"`
	Price     uint32   `json:"price" form:"price" binding:"required"`
	AreaId    uint32   `json:"area_id" form:"area_id" binding:"required"`
	Address   string   `json:"address" form:"address" binding:"required"`
	RoomCount uint32   `json:"room_count" form:"room_count binding:"required""`
	Acreage   uint32   `json:"acreage" form:"acreage" binding:"required"`
	Unit      string   `json:"unit" form:"unit" binding:"required"`
	Capacity  uint32   `json:"capacity" form:"capacity" binding:"required"`
	Beds 	  string   `json:"beds" form:"beds" binding:"required"`
	Deposit   uint32   `json:"deposit" form:"deposit binding:"required"`
	MinDays	  uint32   `json:"min_days" form:"min_days" binding:"required"`
	MaxDays	  uint32   `json:"max_days" form:"max_days""`
	Facility  []uint32 `json:"facility" form:"facility[]""`
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

//房屋信息列表
func UserHouseList(ctx *gin.Context) {
	aid, _ := strconv.Atoi(ctx.Query("aid"))
	sd := ctx.Query("sd")
	ed := ctx.Query("ed")
	page, _ := strconv.Atoi(ctx.Query("p"))

	userInfo, exists := ctx.Get("user_info")
	if !exists {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_SESSIONERR, "请先登录"))
		return
	}
	user := userInfo.(User)
	userId:= user.Id

	request := &house.ListRequest{Aid: uint32(aid), StartDate: sd, EndDate: ed, Page: uint32(page), UserId : userId}

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

	file, err := ctx.FormFile("house_image")
	if err != nil {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_PARAMERR, "图片保存失败1:"+err.Error()))
		return
	}

	fileName := "house_" + houseId + "_" + file.Filename
	path := "/views/images/" + fileName
	err = ctx.SaveUploadedFile(file, "."+path)
	if err != nil {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_SERVERERR, "图片保存失败2:"+err.Error()))
		return
	}

	client := micro.NewService()
	houseService := house.NewHouseService(config.HOUSE, client.Client())

	_, err = houseService.UploadImage(context.TODO(), &house.UploadImageRequest{HouseId: uint32(houseIdInt), Url: "/home/images/"+fileName})
	if err != nil {
		msg := GetServiceError(err.Error())
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_USERERR, "图片保存失败3:"+msg.Detail))
		return
	}

	path = "/home/images/"+fileName

	res := struct{ Url string `json:"url"`}{}
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

	userInfo, exists := ctx.Get("user_info")
	if !exists {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_SESSIONERR, "请先登录"))
		return
	}
	user := userInfo.(User)
	userId:= user.Id

	client := micro.NewService()
	houseService := house.NewHouseService(config.HOUSE, client.Client())

	rsp, err := houseService.Detail(context.TODO(), &house.DetailRequest{HouseId: uint32(houseIdInt)})
	if err != nil {
		msg := GetServiceError(err.Error())
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_SERVERERR, "数据获取失败:"+msg.Detail))
		return
	}
	type result struct {
		House *house.DetailResponse `json:"house"`
		UserId uint32 `json:"user_id"`
	}
	res := result{
		House:  rsp,
		UserId: userId,
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

	userInfo, exists := ctx.Get("user_info")
	if !exists {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_SESSIONERR, "请先登录"))
		return
	}
	user := userInfo.(User)
	userId:= user.Id

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
		UserId: userId,
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
