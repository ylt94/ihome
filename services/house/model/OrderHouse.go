package model

import (
	"time"

	"gorm.io/gorm"
)

type OrderHouse struct {
	gorm.Model            //订单编号
	UserId      uint64    `json:"user_id"`       //下单的用户编号   //与用户表进行关联
	HouseId     uint64    `json:"house_id"`      //预定的房间编号   //与房屋信息进行关联
	HouseUserId uint64    `json:"house_user_id"` //房东用户编号   //与用户表进行关联
	Begin_date  time.Time `gorm:"type:datetime"` //预定的起始时间
	End_date    time.Time `gorm:"type:datetime"` //预定的结束时间
	Days        uint64    //预定总天数
	House_price uint64    //房屋的单价
	Amount      uint64    //订单总金额
	Status      string    `gorm:"default:'WAIT_ACCEPT'"` //订单状态
	Comment     string    `gorm:"size:512"`              //订单评论
	Credit      bool      //表示个人征信情况 true表示良好
}

func (e *OrderHouse) GetDataByHouseIds(datas *[]OrderHouse, houseIds []uint32, fields string, limit uint32, order string) {
	query := Db().Model(e)
	query.Select(fields).Where("house_id in ?", houseIds).Order(order).Limit(int(limit)).Offset(0).Find(datas)
}

func (e *OrderHouse) GetCommentsByHouseId(datas interface{}, houseId uint32, fields string, limit uint32, order string) {
	query := Db().Model(e).Joins("left join user on order_house.user_id = user.id")
	query.Select(fields).Where("house_id = ?", houseId).Order(order).Limit(int(limit)).Offset(0).Find(datas)
}
