package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectDB() (*gorm.DB, error) {
	dsn := "root:123456@tcp(192.168.0.155:3306)/blogV3?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func GetDB() *gorm.DB {
	db, err := connectDB()
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
