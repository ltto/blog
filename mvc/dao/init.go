package dao

import (
	"time"

	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	if db, err := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/hress?charset=utf8mb4&collation=utf8mb4_bin&loc=Local&parseTime=true"); err != nil {
		panic(err)
	} else {
		DB = db
		//defer DB.Close()
		DB.LogMode(true)
		DB.DB().SetConnMaxLifetime(time.Hour)
	}
}
