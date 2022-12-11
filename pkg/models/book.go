package models

import (
	"github.com/mahmudulhasandip/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

// Book model struct
type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// DB connect and migration
func init() {
	config.Connect()
	db = config.GetDatabase()
	if err := db.AutoMigrate(&Book{}); err != nil {
		return
	}
}

// Book DB model func

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	Books := make([]Book, 0)
	db.Find(&Books)
	return Books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("id = ?", id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("id=?", id).Delete(&book)
	return book
}
