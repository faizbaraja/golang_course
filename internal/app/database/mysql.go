package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnDatabase() *gorm.DB {
	//dsn := "root:faizbaraja@tcp(127.0.0.1:3306)/course?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:O7vxgiH40FXMRlQaDexh@tcp(containers-us-west-106.railway.app:6198)/railway?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
