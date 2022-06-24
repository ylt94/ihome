package model

import (
	"time"

	"gorm.io/gorm"
)

/* 订单 table_name = order */
type OrderHouse struct {
	gorm.Model            //订单编号
	UserId      uint32    `json:"user_id"`       //下单的用户编号   //与用户表进行关联
	HouseId     uint32    `json:"house_id"`      //预定的房间编号   //与房屋信息进行关联
	BeginDate   time.Time `gorm:"type:datetime"` //预定的起始时间
	EndDate     time.Time `gorm:"type:datetime"` //预定的结束时间
	Days        uint32    //预定总天数
	HousePrice  uint32    //房屋的单价
	Amount      uint32    //订单总金额
	Status      string    `gorm:"default:'WAIT_ACCEPT'"` //订单状态
	Comment     string    `gorm:"size:512"`              //订单评论
	Credit      bool      //表示个人征信情况 true表示良好
	HouseUserId uint32
}

func (m *OrderHouse) insert() {

}

func (m *OrderHouse) Update(where map[string]WhereItem, updateData map[string]interface{}) {
	query := Db().Model(m)
	query = getWhere(query, where)
	query.Updates(updateData)
}

func (m *OrderHouse) GetOrderByWhere(where map[string]WhereItem) {
	query := Db().Model(m)
	query = getWhere(query, where)
	query.First(m)
}

func (m *OrderHouse) List(data interface{}, where map[string]WhereItem, fields string, order string, page uint32, limit uint32) (uint32, uint32) {
	query := Db().Model(m)
	query = getWhere(query, where)
	query.Select(fields)
	if order != "" {
		query.Order(order)
	}

	var count int64
	if page > 0 {
		countQuery := query
		countQuery.Count(&count)
		offset := (page - 1) * limit
		query.Offset(int(offset)).Limit(int(limit))
	}
	query.Find(data)

	return page, uint32(count)
}
