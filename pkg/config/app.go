package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "host=localhost user=root password=secret dbname=go-bookstore port=5432 sslmode=disable TimeZone=Asia/Dhaka"

//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDatabase() *gorm.DB {
	return db
}
