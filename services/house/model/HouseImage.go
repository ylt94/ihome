package model

/* 房屋图片 table_name = "house_image"*/
type HouseImage struct {
	Id      uint64 `json:"house_image_id"`      //图片id
	Url     string `gorm:"size:256" json:"url"` //图片url     存放我们房屋的图片
	HouseId uint64 `json:"house_id"`            //图片所属房屋编号
}

func (e *HouseImage) GetDataByHouseIds(datas *[]HouseImage, HouseIds []uint32, fields string) {
	query := Db().Model(e)
	query.Select(fields).Where("house_id in ?", HouseIds).Find(datas)
}
