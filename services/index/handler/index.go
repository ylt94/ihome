package handler

import (
	"context"
	_ "encoding/json"

	log "github.com/micro/go-micro/v2/logger"
	models "github.com/ylt94/ihome/services/index/model"

	index "github.com/ylt94/ihome/services/index/proto/index"
)

type Index struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Index) Area(ctx context.Context, req *index.AreaRequest, rsp *index.AreaResponse) error {
	log.Info("Received Index.Call request")
	var res []models.Area
	model := &models.Area{}
	model.GetAreas(&res)

	areas := make([]*index.AreaItem, len(res))
	for _, item := range res {
		areas = append(areas, &index.AreaItem{Id: item.Id, Name: item.Name})
	}

	rsp.Areas = areas
	return nil
}

func (e *Index) Banner(ctx context.Context, req *index.BannerRequest, rsp *index.BannerResponse) error {
	//TODO 请求house service
	var res []models.House
	db := models.Db()
	db.Table("house").Find(&res)

	houses := make([]*index.House, len(res))
	for _, item := range res {
		houses = append(houses, &index.House{
			HouseId: item.Id,
			Title:   item.Title,
		})
	}
	rsp.Houses = houses
	return nil
}
