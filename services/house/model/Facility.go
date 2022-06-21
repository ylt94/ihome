package model

/* 设施信息 table_name = "facility"*/ //设施信息 需要我们提前手动添加的
type Facility struct {
	Id     uint64   `json:"fid"`     //设施编号
	Name   string   `gorm:"size:32"` //设施名字
	Houses []*House //都有哪些房屋有此设施  与房屋表进行关联的
}

func (e *Facility) GetDataByHouseIds(datas *[]HouseImage, HouseIds []uint) {
	query := Db().Model(e)
	query.Where("house_id in ?", HouseIds).Find(datas)
}
