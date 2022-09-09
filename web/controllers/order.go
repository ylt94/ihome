package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	order "github.com/ylt94/ihome/services/order/proto/order"
	"github.com/ylt94/ihome/web/config"
	"github.com/ylt94/ihome/web/utils"
)

//租房下单
func OrderBuy(ctx *gin.Context) {

	request := &order.CreateRequest{}

	houseId := ctx.PostForm("house_id")
	houseIdInt, err := strconv.Atoi(houseId)
	if err != nil {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_PARAMERR, "参数错误:house_id"+err.Error()))
		return
	}

	if houseIdInt == 0 {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_PARAMERR, "缺少参数错误:house_id"))
		return
	}
	request.HouseId = uint32(houseIdInt)

	startDate := ctx.PostForm("start_date")
	if startDate == "" {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_PARAMERR, "请选择起始日期"))
		return
	}
	//TODO 验证日期格式
	request.StartDate = startDate

	endDate := ctx.PostForm("end_date")
	if endDate == "" {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_PARAMERR, "请选择结束日期"))
		return
	}
	//TODO 验证日期格式
	request.EndDate = endDate

	userInfo, exists := ctx.Get("user_info")
	if !exists {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_SESSIONERR, "请先登录"))
		return
	}
	user := userInfo.(User)
	request.UserId = user.Id

	client := micro.NewService()
	orderService := order.NewOrderService(config.ORDER, client.Client())

	rsp, err := orderService.Create(context.TODO(), request)
	if err != nil {
		msg := GetServiceError(err.Error())
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_SERVERERR, "下单失败:"+msg.Detail))
		return
	}

	ctx.JSON(http.StatusOK, GetReturn(rsp, utils.RECODE_OK, "下单成功"))
	return
}

//订单列表
func OrderList(ctx *gin.Context) {
	role := ctx.Query("role")
	if role == "" {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_PARAMERR, "缺少参数:role"))
		return
	}

	userInfo, exists := ctx.Get("user_info")
	if !exists {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_SESSIONERR, "请先登录"))
		return
	}
	user := userInfo.(User)
	UserId := user.Id

	client := micro.NewService()
	orderService := order.NewOrderService(config.ORDER, client.Client())

	rsp, err := orderService.List(context.TODO(), &order.ListRequest{Role: role, UserId: UserId})
	if err != nil {
		msg := GetServiceError(err.Error())
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_SERVERERR, "请求失败:"+msg.Detail))
		return
	}
	ctx.JSON(http.StatusOK, GetReturn(rsp, utils.RECODE_OK, "成功:"))
	return
}

//更新订单状态
func OrderUpdate(ctx *gin.Context) {
	orderId := ctx.PostForm("order_id")
	status := ctx.PostForm("status")

	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_PARAMERR, "参数错误:order_id"+err.Error()))
		return
	}
	if orderIdInt == 0 {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_PARAMERR, "缺少参数:order_id"))
		return
	}

	statusInt, ok := order.OrderStatus_value[status]
	if !ok {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_PARAMERR, "缺少参数:status"))
		return
	}

	userInfo, exists := ctx.Get("user_info")
	if !exists {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_SESSIONERR, "请先登录"))
		return
	}
	user := userInfo.(User)
	UserId := user.Id

	client := micro.NewService()
	orderService := order.NewOrderService(config.ORDER, client.Client())

	_, err = orderService.UpdateStatus(context.TODO(), &order.StatusRequest{Status: order.OrderStatus(statusInt), UserId: UserId, OrderId: uint32(orderIdInt)})
	if err != nil {
		msg := GetServiceError(err.Error())
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_SERVERERR, "操作失败:"+msg.Detail))
		return
	}

	ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_OK, "操作成功"))
	return
}

//评论
func OrderComment(ctx *gin.Context) {
	orderId := ctx.PostForm("order_id")
	comment := ctx.PostForm("comment")

	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_PARAMERR, "参数错误:order_id"+err.Error()))
		return
	}
	if orderIdInt == 0 {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_PARAMERR, "缺少参数:order_id"))
		return
	}
	if comment == "" {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_PARAMERR, "请输入你的评论"))
		return
	}

	userInfo, exists := ctx.Get("user_info")
	if !exists {
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_SESSIONERR, "请先登录"))
		return
	}
	user := userInfo.(User)
	UserId := user.Id

	client := micro.NewService()
	orderService := order.NewOrderService(config.ORDER, client.Client())

	_, err = orderService.Comment(context.TODO(), &order.CommentRequest{OrderId: uint32(orderIdInt), UserId: UserId, Comment: comment})
	if err != nil {
		msg := GetServiceError(err.Error())
		ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_SERVERERR, "评论失败:"+msg.Detail))
		return
	}

	ctx.JSON(http.StatusOK, GetReturn("", utils.RECODE_OK, "评论成功"))
	return
}
