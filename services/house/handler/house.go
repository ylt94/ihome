package handler

import (
	"context"

	"github.com/ylt94/ihome/services/house/model"

	log "github.com/micro/go-micro/v2/logger"

	house "github.com/ylt94/ihome/services/house/proto/house"
)

type House struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *House) List(ctx context.Context, req *house.ListRequest, rsp *house.ListResponse) error {
	log.Info("Received House.List request")
	where := make(map[string]model.WhereItem, 4)
	if req.GetAid() > 0 {
		where["area_id"] = model.WhereItem{Condition: "=", Val: req.GetAid()}
	}
	if req.GetStartDate() != "" {
		where["start_date"] = model.WhereItem{Condition: ">=", Val: req.GetStartDate()}
	}
	if req.GetEndDate() != "" {
		where["end_date"] = model.WhereItem{Condition: "<=", Val: req.GetEndDate()}
	}

	//获取house表数据
	offset := 10
	page := req.GetPage()
	fields := "address,area_id,ctime,id,index_image_url,order_count,price,room_count,title,user_id"

	mainData := make([]model.House, 0, offset)
	house := &model.House{}
	_, count := house.GetList(&mainData, where, fields, int(page), offset)
	rsp.CurrentPage = page
	rsp.TotalPage = uint32(count)

	areaIds := make([]uint32, 0, offset)
	userIds := make([]uint32, 0, offset)
	for k, _ := range mainData {
		areaIds = append(areaIds, mainData[k].AreaId)
		userIds = append(userIds, mainData[k].UserId)
	}

	//获取用户信息
	userData := make([]model.User, 0, offset)
	(&model.User{}).GetUserByIds(userData, userIds, "id, user_avatar")

	userMapData := make([uint32]string, offset)
	for k, _ := range userData {
		userMapData[userData[k].Id] = userData[k].AvatarUrl
	}

	//获取地区信息

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
