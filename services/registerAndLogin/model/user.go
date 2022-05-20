package model

type User struct {
	Id int
	Name string
	PasswordHash string
	Mobile string
	RealName string
	IdCard string
	AvatarUrl string
}

func (u *User) Register() {
	db := Db()
	db.Create(u)
}
