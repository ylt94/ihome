package test

import (
	"context"
	"log"
	"testing"

	"github.com/ylt94/ihome/services/order/handler"

	order "github.com/ylt94/ihome/services/order/proto/order"
)

func TestCreate(T *testing.T) {
	req := &order.CreateRequest{
		HouseId:   1,
		UserId:    2,
		StartDate: "2022-09-07",
		EndDate:   "2022-09-11",
	}
	rsp := &order.CreateResponse{}

	err := new(handler.Order).Create(context.TODO(), req, rsp)
	if err != nil {
		log.Println(err)
	}
	log.Println(rsp)
}

func TestList(T *testing.T) {
	req := &order.ListRequest{
		Role:   "houser",
		UserId: 1,
	}
	rsp := &order.ListResponse{}

	err := new(handler.Order).List(context.TODO(), req, rsp)
	if err != nil {
		log.Println(err)
	}
	log.Println(rsp)
}
