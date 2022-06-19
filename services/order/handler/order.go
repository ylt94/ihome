package handler

import (
	"context"

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
	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Order) Comment(ctx context.Context, req *order.CommentRequest, rsp *order.CommentResponse) error {
	log.Info("Received Order.Call request")
	return nil
}
