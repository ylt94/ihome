package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var DB *gorm.DB

func Db() *gorm.DB {
	if DB != nil {
		return DB
	}
	dsn := "root:123456@tcp(127.0.0.1:3306)/ihome?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		panic(err.Error())
	}

	ConnPool, err := DB.DB()
	if err != nil {
		panic(err.Error())
	}
	ConnPool.SetConnMaxLifetime(60*time.Second)
	ConnPool.SetMaxIdleConns(5)
	ConnPool.SetMaxOpenConns(10)
	return DB
}