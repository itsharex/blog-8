package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDb() {
	db, err = gorm.Open(mysql.Open("root:1234abc.@(127.0.0.1:3306)/ginblog?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数:", err)

	}
	db.AutoMigrate(&User{}, &Article{}, &Category{})

}
