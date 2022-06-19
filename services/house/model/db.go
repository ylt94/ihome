package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type  WhereItem struct {
	Condition string
	Val interface{}
}

var Conn *gorm.DB

func Db() *gorm.DB {
	if Conn != nil {
		return Conn
	}
	dsn := "root:123456@tcp(127.0.0.1:3306)/ihome?charset=uft8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		panic(err.Error())
	}
	Conn = db
	return Conn
}


func getWhere(query *gorm.DB, where map[string]WhereItem) *gorm.DB {
	for c,_ := range where {
		v := where[c]
		whereStr := c+" "+v.Condition + " ?"
		query.Where(whereStr, v.Val)
	}
	return query
}
