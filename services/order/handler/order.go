package handler

import (
	"context"
	"errors"
	"github.com/ylt94/ihome/services/order/model"

	log "github.com/micro/go-micro/v2/logger"

	order "github.com/ylt94/ihome/services/order/proto/order"
)

type Order struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Order) Create(ctx context.Context, req *order.CreateRequest, rsp *order.CreateResponse) error {
	log.Info("Received Order.create request")
	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Order) List(ctx context.Context, req *order.ListRequest, rsp *order.ListResponse) error {
	log.Info("Received Order.List request")
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

	orderModel := new(model.OrderHouse)

	where := make(map[string]model.WhereItem, 3)
	where["id"] = model.WhereItem{"=", req.OrderId}
	where["house_user_id"] = model.WhereItem{"=", req.UserId}
	where["status"] = model.WhereItem{"=", order.OrderStatus_WAIT}

	orderModel.GetOrderByWhere(where)
	if orderModel.ID == 0 {
		return errors.New("订单不存在")
	}

	updateData := map[string]interface{}{"status" : req.Status}
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
	where["id"] = model.WhereItem{"=", req.OrderId}
	where["user_id"] = model.WhereItem{"=", req.UserId}

	orderModel.GetOrderByWhere(where)
	if orderModel.ID == 0 {
		return errors.New("订单不存在")
	}

	updateData := map[string]interface{}{"comment" : req.Comment}
	orderModel.Update(where, updateData)

	return nil
}
