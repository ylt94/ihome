package model

/* 设施信息 table_name = "facility"*/ //设施信息 需要我们提前手动添加的
type HouseFacility struct {
	HouseId      uint32
	FacilityId   uint32
}

type Facility struct {
	Id   uint32
	Name string
}

func (e *HouseFacility) GetHouseDataByHouseIds(datas interface{}, HouseIds []uint32) {
	query := Db().Model(e)
	query.Where("house_id in ?", HouseIds)
	query.Select("house_facilities.house_id, facility.name")
	query.Joins("left join facility on facility.id = house_facilities.house_id").Find(datas)
}
