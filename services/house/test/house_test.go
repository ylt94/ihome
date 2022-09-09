package test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/ylt94/ihome/services/house/handler"

	house "github.com/ylt94/ihome/services/house/proto/house"
)

func TestList(T *testing.T) {
	req := &house.ListRequest{}
	rsp := &house.ListResponse{}
	obj := &handler.House{}
	err := obj.List(context.TODO(), req, rsp)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(rsp.GetHouses())
}

func TestCreate(T *testing.T) {
	req := &house.CreateRequest{
		Title:     "测试添加标题",
		Price:     40000,
		AreaId:    1,
		Address:   "测试添加地址",
		RoomCount: 1,
		Acreage:   20,
		Unit:      "一间",
		Capacity:  1,
		Beds:      "1张",
		Deposit:   40000,
		MinDays:   1,
		MaxDays:   0,
		Facility:  []uint32{1},
		UserId: 1,
	}

	rsp := &house.CreateResponse{}
	obj := &handler.House{}
	err := obj.Create(context.TODO(), req, rsp)
	if err != nil {
		log.Println("错误:", err)
	}
	fmt.Println(rsp)
}

func TestDetail(T *testing.T) {
	req := &house.DetailRequest{HouseId: 1}
	rsp := &house.DetailResponse{}
	obj := &handler.House{}
	err := obj.Detail(context.TODO(), req, rsp)
	if err != nil {
		log.Println(err)
	}
	log.Println(rsp)
}

func TestUploadImage(T *testing.T) {
	req := &house.UploadImageRequest{HouseId: 3, Url: "/home/images/house_3_house_test2.jpg"}
	rsp := &house.UploadImageResponse{}
	obj := &handler.House{}
	err := obj.UploadImage(context.TODO(), req, rsp)
	if err != nil {
		log.Println(err)
	}
	log.Println(rsp)
}
