package handler

import (
	"context"
	_ "encoding/json"
	models "github.com/ylt94/ihome/services/index/model"

	log "github.com/micro/go-micro/v2/logger"

	index "github.com/ylt94/ihome/services/index/proto/index"
)

type Index struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Index) Area(ctx context.Context, req *index.AreaRequest, rsp *index.AreaResponse) error {
	log.Info("Received Index.Call request")
	//var res []models.Area
	//model := &models.Area{}
	//model.GetAreas(&res)
	//dataJson,err := json.Marshal(res)
	//if err != nil {
	//	return err
	//}
	//
	//var areas []*index.AreaItem
	//json.Unmarshal(dataJson, &areas)

	var areas []*index.AreaItem
	db := models.Db()
	db.Table("area").Find(&areas)

	rsp.Areas = areas
	return nil
}

func (e *Index) Banner(ctx context.Context, req *index.BannerRequest, rsp *index.BannerResponse) error {
	var houses []*index.House
	db := models.Db()
	db.Table("house").Find(&houses)
	rsp.Houses = houses
	return nil
}
