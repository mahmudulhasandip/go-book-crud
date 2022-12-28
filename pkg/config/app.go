package config

import (
	"github.com/davecgh/go-spew/spew"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

//var dsn = "host=localhost user=root password=secret dbname=go-bookstore port=5432 sslmode=disable TimeZone=Asia/Dhaka"

//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

var (
	db *gorm.DB
)

func Connect() {
	LoadEnv()
	dsn := os.Getenv("DB")
	spew.Dump(dsn)
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDatabase() *gorm.DB {
	return db
}
