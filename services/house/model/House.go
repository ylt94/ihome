package model

import "math"

type House struct {
	Id            uint32 //房屋编号
	UserId        uint32 //房屋主人的用户编号  与用户进行关联
	AreaId        uint32 //归属地的区域编号   和地区表进行关联
	Title         string `gorm:"size:64" `                 //房屋标题
	Address       string `gorm:"size:512"`                 //地址
	RoomCount     uint32 `gorm:"default:1" `               //房间数目
	Acreage       uint32 `gorm:"default:0" json:"acreage"` //房屋总面积
	Price         uint32 `json:"price"`
	Unit          string `gorm:"size:32;default:''" json:"unit"`             //房屋单元,如 几室几厅
	Capacity      uint32 `gorm:"default:1" json:"capacity"`                  //房屋容纳的总人数
	Beds          string `gorm:"size:64;default:''" json:"beds"`             //房屋床铺的配置
	Deposit       uint32 `gorm:"default:0" json:"deposit"`                   //押金
	MinDays       uint32 `gorm:"default:1" json:"min_days"`                  //最少入住的天数
	MaxDays       uint32 `gorm:"default:0" json:"max_days"`                  //最多入住的天数 0表示不限制
	OrderCount    uint32 `gorm:"default:0" json:"order_count"`               //预定完成的该房屋的订单数
	IndexImageUrl string `gorm:"size:256;default:''" json:"index_image_url"` //房屋主图片路径
	CreatedAt     string
	UpdatedAt     string
	DeletedAt     string
}

func (e *House) GetList(data interface{}, where map[string]WhereItem, fields string, page int, limit int) (uint32, uint32) {
	db := Db()
	query := db.Model(e)

	query = getWhere(query, where)
	countQuery := query

	query.Select(fields).Joins("left join area on house.area_id = area.id")

	var count int64
	if page > 0 {
		countQuery.Count(&count)
		offset := (page - 1) * limit
		query.Limit(limit).Offset(offset)
	}

	query.Find(data)

	totalPage := math.Ceil(float64(count)/float64(limit))
	return uint32(page), uint32(totalPage)
}

func (e *House) Detail(HouseId uint32, fields string) {
	db := Db()
	db.Select(fields).Where("id = ?", HouseId)
	db.First(e)
}

func (e *House) InsertIndexImage (houseId uint32, updateData map[string]interface{}) {
	Db().Model(e).Where("id = ?", houseId).Updates(updateData)
}

func (e *House) Create() {
	Db().Create(e)
}