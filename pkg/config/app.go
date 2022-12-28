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
	//dsn := "host=containers-us-west-178.railway.app user=postgres password=y9Ethx9TbT2fLzAb3Iaj dbname=railway port=7372 sslmode=disable TimeZone=Asia/Dhaka"
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
