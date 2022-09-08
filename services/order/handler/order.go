package handler

import (
	"context"
	"errors"
	"math"
	"time"

	"github.com/ylt94/ihome/services/order/model"

	log "github.com/micro/go-micro/v2/logger"

	order "github.com/ylt94/ihome/services/order/proto/order"
)

const TIMEFORMAT = "2006-01-02"

type Order struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Order) Create(ctx context.Context, req *order.CreateRequest, rsp *order.CreateResponse) error {
	log.Info("Received Order.create request")

	//查询房屋信息
	houseModel := new(model.House)
	where := make(map[string]model.WhereItem, 1)
	where["id"] = model.WhereItem{Condition: "=", Val: req.HouseId}
	houseModel.GetOne(where, "*")
	if houseModel.Id == 0 {
		return errors.New("房屋信息已不存在或已下架")
	}

	//根据开始时间和结束时间计算天数
	startTime, err := time.ParseInLocation(TIMEFORMAT, req.StartDate, time.Local)
	if err != nil {
		return err
	}

	endTime, err := time.ParseInLocation(TIMEFORMAT, req.EndDate, time.Local)
	if err != nil {
		return err
	}
	days := uint32(math.Ceil(endTime.Sub(startTime).Hours() / 24))

	OrderModel := model.OrderHouse{
		UserId:      req.UserId,
		HouseId:     houseModel.Id,
		BeginDate:   req.StartDate,
		EndDate:     req.EndDate,
		Days:        days,
		HousePrice:  houseModel.Price,
		Amount:      houseModel.Price * days,
		Status:      order.OrderStatus_WAIT.String(),
		Credit:      1,
		HouseUserId: houseModel.UserId,
	}
	OrderModel.Create()
	if OrderModel.ID == 0 {
		return errors.New("订单创建失败")
	}
	rsp.OrderId = OrderModel.ID
	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Order) List(ctx context.Context, req *order.ListRequest, rsp *order.ListResponse) error {
	log.Info("Received Order.List request")
	user_id_key := "user_id"
	if req.Role == "houser" {
		user_id_key = "house_user_id"
	}
	where := make(map[string]model.WhereItem, 1)
	where[user_id_key] = model.WhereItem{Condition: "=", Val: req.UserId}

	//获取订单主数据
	data := []model.OrderHouse{}
	fileds := "amount, comment, created_at, days, end_date, id, begin_date, status, house_id"
	new(model.OrderHouse).List(&data, where, fileds, "id desc", 0, 0)
	if len(data) == 0 {
		return nil
	}

	houseIds := make([]uint32, 0, len(data))
	for _, v := range data {
		houseIds = append(houseIds, v.HouseId)
	}

	//请求house 服务获取房屋信息--暂时直接查
	houseWhere := make(map[string]model.WhereItem)
	houseWhere["id"] = model.WhereItem{Condition: "in", Val: houseIds}
	houseData := make([]model.House, 0, len(houseIds))
	new(model.House).GetList(&houseData, houseWhere, "id, index_image_url, title")
	houseMapData := make(map[uint32]model.House, len(houseData))
	for _, v := range houseData {
		houseMapData[v.Id] = model.House{IndexImageUrl: v.IndexImageUrl, Title: v.Title}
	}

	//组装返回数据
	for _, v := range data {
		item := order.ListItem{
			Amount:    v.Amount,
			Comment:   v.Comment,
			Ctime:     v.CreatedAt.Format(TIMEFORMAT),
			Days:      v.Days,
			EndDate:   v.EndDate,
			ImgUrl:    houseMapData[v.HouseId].IndexImageUrl,
			OrderId:   v.ID,
			StartDate: v.BeginDate,
			Status:    v.Status,
			Title:     houseMapData[v.HouseId].Title,
		}
		rsp.Orders = append(rsp.Orders, &item)
	}
	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Order) UpdateStatus(ctx context.Context, req *order.StatusRequest, rsp *order.StatusResponse) error {
	log.Info("Received Order.UpdateStatus request")
	if req.UserId == 0 || req.OrderId == 0 {
		return errors.New("缺少参数")
	}
	if req.Status != order.OrderStatus_ACCEPT && req.Status != order.OrderStatus_REJECT {
		return errors.New("错误的订单状态")
	}
	status := req.Status.String()
	orderModel := new(model.OrderHouse)

	where := make(map[string]model.WhereItem, 3)
	where["id"] = model.WhereItem{Condition: "=", Val: req.OrderId}
	where["house_user_id"] = model.WhereItem{Condition: "=", Val: req.UserId}
	where["status"] = model.WhereItem{Condition: "=", Val: order.OrderStatus_WAIT}

	orderModel.GetOrderByWhere(where)
	if orderModel.ID == 0 {
		return errors.New("订单不存在")
	}

	updateData := map[string]interface{}{"status": status}
	orderModel.Update(where, updateData)

	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Order) Comment(ctx context.Context, req *order.CommentRequest, rsp *order.CommentResponse) error {
	log.Info("Received Order.Call request")
	if req.UserId == 0 || req.OrderId == 0 {
		return errors.New("缺少参数")
	}

	orderModel := new(model.OrderHouse)

	where := make(map[string]model.WhereItem, 2)
	where["id"] = model.WhereItem{Condition: "=", Val: req.OrderId}
	where["user_id"] = model.WhereItem{Condition: "=", Val: req.UserId}

	orderModel.GetOrderByWhere(where)
	if orderModel.ID == 0 {
		return errors.New("订单不存在")
	}

	updateData := map[string]interface{}{"comment": req.Comment}
	orderModel.Update(where, updateData)

	return nil
}
