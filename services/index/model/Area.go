package model

type Area struct {
	Id   uint64
	Name string
}

func (this *Area) GetAreas(data *[]Area) *[]Area {
	db := Db()
	db.Find(data)
	return data
}
