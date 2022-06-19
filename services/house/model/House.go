package model

import "gorm.io/gorm"

type House struct {
	gorm.Model                     	 //房屋编号
	UserId          uint64           //房屋主人的用户编号  与用户进行关联
	AreaId          uint64           //归属地的区域编号   和地区表进行关联
	Title           string        	 `gorm:"size:64" `         //房屋标题
	Address         string        	 `gorm:"size:512"`         //地址
	RoomCount       uint64           `gorm:"default:1" `       //房间数目
	Acreage         uint64           `gorm:"default:0" json:"acreage"`                   //房屋总面积
	Price           uint64           `json:"price"`
	Unit            string           `gorm:"size:32;default:''" json:"unit"`             //房屋单元,如 几室几厅
	Capacity        uint64           `gorm:"default:1" json:"capacity"`                   //房屋容纳的总人数
	Beds            string           `gorm:"size:64;default:''" json:"beds"`             //房屋床铺的配置
	Deposit         uint64           `gorm:"default:0" json:"deposit"`                   //押金
	MinDays         uint64           `gorm:"default:1" json:"min_days"`                   //最少入住的天数
	MaxDays         uint64           `gorm:"default:0" json:"max_days"`                   //最多入住的天数 0表示不限制
	OrderCount      uint64           `gorm:"default:0" json:"order_count"`               //预定完成的该房屋的订单数
	IndexImageUrl   string           `gorm:"size:256;default:''" json:"index_image_url"` 			 //房屋主图片路径
	Facilities      []*Facility      `gorm:"many2many:house_facilities" json:"facilities"`                 //房屋设施   与设施表进行关联
	Images          []*HouseImage    `json:"img_urls"`                      //房屋的图片   除主要图片之外的其他图片地址
	Orders          []*OrderHouse    `json:"orders"`                        //房屋的订单    与房屋表进行管理
}

func(e *House) getList(data *[]House, where map[string]WhereItem, fields string, page int, limit int) (*[]House, uint64,uint64){
	db := Db()
	query := db.Model(e)

	query = getWhere(query, where)
	countQuery := query

	var count int64
	countQuery.Count(&count)

	offset := (page-1)*limit
	query.Select(fields).Limit(limit).Offset(offset).Find(data)

	return data, uint64(page), uint64(count)
}