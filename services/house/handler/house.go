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
		where["house.area_id"] = model.WhereItem{Condition: "=", Val: req.GetAid()}
	}
	if req.GetStartDate() != "" {
		where["house.created_at"] = model.WhereItem{Condition: ">=", Val: req.GetStartDate()}
	}
	if req.GetEndDate() != "" {
		where["house.created_at"] = model.WhereItem{Condition: "<=", Val: req.GetEndDate()}
	}

	//获取house表数据
	offset := 10
	page := req.GetPage()
	fields := "house.address,house.area_id,house.created_at," +
		"house.id,house.index_image_url,house.order_count," +
		"house.price,house.room_count,house.title,house.user_id,area.name as area_name"
	type list struct {
		Address       string
		AreaId        uint32
		CreatedAt     string
		Id            uint32
		IndexImageUrl string
		OrderCount    uint32
		Price         uint32
		RoomCount     uint32
		Title         string
		UserId        uint32
		AreaName      string
	}
	mainData := make([]list, 0, offset)


	houseModel := &model.House{}
	_, count := houseModel.GetList(&mainData, where, fields, int(page), offset)

	rsp.CurrentPage = page
	rsp.TotalPage = uint32(count)

	userIds := make([]uint32, 0, offset)
	houseIds := make([]uint32, 0, offset)
	for k, _ := range mainData {
		userIds = append(userIds, mainData[k].UserId)
		houseIds = append(houseIds, mainData[k].Id)
	}

	//获取用户信息
	userData := make([]model.User, 0, offset)
	(&model.User{}).GetUserByIds(&userData, userIds, "id, user_avatar")

	userMapData := make(map[uint32]string, offset)
	for k, _ := range userData {
		userMapData[userData[k].Id] = userData[k].AvatarUrl
	}

	for k, _ := range mainData {
		item := mainData[k]
		listItem := house.ListItem{
			Address:item.Address,
			AreaName:item.AreaName,
			Ctime:item.CreatedAt,
			HouseId:item.Id,
			ImageUrl:item.IndexImageUrl,
			OrderCount:item.OrderCount,
			Price:item.Price,
			RoomCount:item.RoomCount,
			Title:item.Title,
			UserAvatar:userMapData[item.Id],
		}
		rsp.Houses = append(rsp.Houses, &listItem)
	}

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
