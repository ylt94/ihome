package model

type User struct {
	Id           uint32
	Name         string
	PasswordHash string
	Mobile       string
	RealName     string
	IdCard       string
	AvatarUrl    string
}

func (e *User) GetUserByIds(users *[]User, userIds []uint32, fields string) {
	db := Db()
	db.Select(fields).Where("id in ?", userIds).Find(users)
}
