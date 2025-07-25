package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	MySQLDSN = "root:csc20040312@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(MySQLDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	return db
}
