package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	house "github.com/ylt94/ihome/services/house/proto/house"
)

type House struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *House) List(ctx context.Context, req *house.ListRequest, rsp *house.ListResponse) error {
	log.Info("Received House.List request")

	return nil
}

func (e *House) Create(ctx context.Context, req *house.CreateRequest, rsp *house.CreateResponse) error {
	log.Info("Received House.Create request")
	return nil
}

func (e *House) UploadImage(ctx context.Context, req *house.UploadImageRequest, rsp *house.UploadImageResponse) error {
	log.Info("Received House.UploadImage request")
	return nil
}


func (e *House) Detail(ctx context.Context, req *house.DetailRequest, rsp *house.DetailResponse) error {
	log.Info("Received House.Detail request")
	return nil
}
