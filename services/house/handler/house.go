package handler

import (
	"context"
	"errors"
	"fmt"
	"sync"

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

	if req.UserId != 0 {
		where["house.user_id"] = model.WhereItem{Condition: "=", Val: req.GetUserId()}
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

	houseModel := new(model.House)
	_, count := houseModel.GetList(&mainData, where, fields, int(page), offset)

	rsp.CurrentPage = page
	rsp.TotalPage = count
	if count == 0 {
		return nil
	}

	userIds := make([]uint32, 0, offset)
	for _, item := range mainData {
		userIds = append(userIds, item.UserId)
	}

	//获取用户信息
	userData := make([]model.User, 0, offset)
	new(model.User).GetUserByIds(&userData, userIds, "id, user_avatar")

	userMapData := make(map[uint32]string, offset)
	for _, item := range userData {
		userMapData[item.Id] = item.AvatarUrl
	}

	for _, item := range mainData {
		listItem := house.ListItem{
			Address:    item.Address,
			AreaName:   item.AreaName,
			Ctime:      item.CreatedAt,
			HouseId:    item.Id,
			ImageUrl:   item.IndexImageUrl,
			OrderCount: item.OrderCount,
			Price:      item.Price,
			RoomCount:  item.RoomCount,
			Title:      item.Title,
			UserAvatar: userMapData[item.Id],
		}
		rsp.Houses = append(rsp.Houses, &listItem)
	}

	return nil
}

func (e *House) Create(ctx context.Context, req *house.CreateRequest, rsp *house.CreateResponse) error {
	log.Info("Received House.Create request")
	houseModel := new(model.House)
	houseModel.Title = req.Title
	houseModel.Price = req.Price
	houseModel.AreaId = req.AreaId
	houseModel.Address = req.Address
	houseModel.RoomCount = req.RoomCount
	houseModel.Acreage = req.Acreage
	houseModel.Unit = req.Unit
	houseModel.Capacity = req.Capacity
	houseModel.Beds = req.Beds
	houseModel.Deposit = req.Deposit
	houseModel.MinDays = req.MinDays
	houseModel.MaxDays = req.MaxDays
	houseModel.Create()

	tx := model.Db().Begin()
	if err := tx.Create(houseModel).Error; err != nil {
		tx.Rollback()
		return err
	}

	var houseId uint32
	rsp.HouseId, houseId = houseModel.Id, houseModel.Id
	facilityIds := req.Facility
	insertData := make([]model.HouseFacility, 0, len(facilityIds))
	for _, v := range facilityIds {
		insertData = append(insertData, model.HouseFacility{HouseId: houseId, FacilityId: v})
	}

	if err := tx.Create(insertData).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (e *House) UploadImage(ctx context.Context, req *house.UploadImageRequest, rsp *house.UploadImageResponse) error {
	log.Info("Received House.UploadImage request")
	houseId := req.HouseId
	url := req.Url
	houseModel := new(model.House)
	houseModel.Detail(houseId, "index_image_url")
	if len(houseModel.IndexImageUrl) == 0 {
		houseModel.InsertIndexImage(houseId, map[string]interface{}{"index_image_url": url})
	} else {
		imageModel := new(model.HouseImage)
		imageModel.InsertImages(houseId, []string{url})
	}
	rsp.Url = url
	return nil
}

func (e *House) Detail(ctx context.Context, req *house.DetailRequest, rsp *house.DetailResponse) error {
	log.Info("Received House.Detail request")
	//房屋主信息
	houseId := req.GetHouseId()
	if houseId == 0 {
		return errors.New("缺少参数house_id")
	}

	houses := model.House{}
	(&houses).Detail(houseId, "*")
	if houses.Id == 0 {
		return fmt.Errorf("房屋信息不存在或已被删除!")
	}

	//协程获取
	wg := sync.WaitGroup{}
	wg.Add(3)

	//查询图片信息
	cap := 10
	images := make([]string, 0, cap)
	go func(images *[]string, cap int) {
		defer wg.Done()
		imgData := make([]model.HouseImage, 0, cap)
		new(model.HouseImage).GetDataByHouseIds(&imgData, []uint32{houseId}, "house_id,url")
		for _, v := range imgData {
			*images = append(*images, v.Url)
		}
	}(&images, cap)

	//获取标签信息
	cap = 25
	facilitys := make([]string, 0, cap)
	go func(facilitys *[]string, cap int) {
		defer wg.Done()
		facilityData := make([]struct{ Name string }, 0, cap)
		new(model.HouseFacility).GetHouseDataByHouseIds(facilityData, []uint32{houseId}, "house_facilities.house_id, facility.name")
		for _, v := range facilityData {
			*facilitys = append(*facilitys, v.Name)
		}
	}(&facilitys, cap)

	//获取订单评论
	var limit uint32 = 25
	comments := make([]*house.Comment, 0, limit)
	go func(comments *[]*house.Comment, limit uint32) {
		defer wg.Done()
		orderData := make([]struct {
			Name      string
			Comment   string
			CreatedAt string
		}, 0, limit)
		new(model.OrderHouse).GetCommentsByHouseId(&orderData, houseId, "user.name, order_house.comment,order_house.created_at", limit, "order_house.id desc")
		for _, item := range orderData {
			comment := house.Comment{
				UserName: item.Name,
				Comment:  item.Comment,
				Ctime:    item.CreatedAt,
			}
			*comments = append(*comments, &comment)
		}
	}(&comments, limit)

	//获取用户信息
	user := new(model.User)
	go func(user *model.User) {
		defer wg.Done()
		user.GetUserById(houses.UserId, "user_id, name, avatar_url")
	}(user)
	wg.Wait()

	//组装返回信息
	rsp.Acreage = houses.Acreage
	rsp.Address = houses.Address
	rsp.Beds = houses.Beds
	rsp.Capacity = houses.Capacity
	rsp.Deposit = houses.Deposit
	rsp.Hid = houses.Id
	rsp.MinDays = houses.MinDays
	rsp.MaxDays = houses.MaxDays
	rsp.Price = houses.Price
	rsp.RoomCount = houses.RoomCount
	rsp.Title = houses.Title
	rsp.Unit = houses.Unit
	rsp.UserId = houses.UserId
	rsp.UserName = user.Name
	rsp.UserAvatar = user.AvatarUrl
	rsp.Comments = comments
	rsp.ImgUrls = images
	rsp.Facilities = facilitys

	return nil
}
