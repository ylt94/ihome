package model

/* 设施信息 table_name = "facility"*/ //设施信息 需要我们提前手动添加的
type HouseFacility struct {
	HouseId    uint32
	FacilityId uint32
}

type Facility struct {
	Id   uint32
	Name string
}

func (e *HouseFacility) TableName() string {
	return "house_facilities"
}

func (e *HouseFacility) GetHouseDataByHouseIds(datas interface{}, HouseIds []uint32, fields string) {
	query := Db().Model(e)
	query.Where("house_id in ?", HouseIds)
	query.Select(fields)
	query.Joins("left join facility on facility.id = house_facilities.facility_id").Find(datas)
}

func (e *HouseFacility) Insert(houseId uint32, facilityIds []uint32) {
	insertData := make([]HouseFacility, 0, len(facilityIds))

	for _, v := range facilityIds {
		insertData = append(insertData, HouseFacility{HouseId: houseId, FacilityId: v})
	}
	Db().Model(e).Create(&insertData)
}
