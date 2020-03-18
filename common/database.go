package common

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"irisProject/model"
)

//原始数据库引擎
var DB *gorm.DB

//启动并连接数据库
func InitDbEngine() *gorm.DB {
	db, err := gorm.Open("mysql", "root:adminwss@/ginessential?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database ,err: " + err.Error())
	}

	DB = db

	InitTable()

	return db
}

func InitTable() {
	DB.AutoMigrate(&model.User{}) //创建user表
}

//对外使用的数据库引擎
func GetDbEngine() *gorm.DB {
	return DB
}
