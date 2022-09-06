package test

import (
	"context"
	"log"
	"testing"

	handler "github.com/ylt94/ihome/services/index/handler"
	index "github.com/ylt94/ihome/services/index/proto/index"
)

func TestArea(T *testing.T) {
	ctx := context.TODO()
	req := &index.AreaRequest{}
	rsp := &index.AreaResponse{}

	obj := &handler.Index{}
	err := obj.Area(ctx, req, rsp)
	log.Println(*rsp)
	if err != nil {
		log.Println(err)
	}
}

func TestBanner(T *testing.T) {
	ctx := context.TODO()
	req := &index.BannerRequest{}
	rsp := &index.BannerResponse{}

	obj := handler.Index{}
	err := obj.Banner(ctx, req, rsp)
	log.Println(*rsp)
	if err != nil {
		log.Println(err)
	}
}
