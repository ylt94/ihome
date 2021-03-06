package model

import "errors"

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

func (u *User) GetUserByPhone(phone string) error {
	if phone == "" {
		return errors.New("请输入手机号!")
	}

	db:= Db()
	db.Where("mobile = ? ", phone).First(u)
	return nil
}
