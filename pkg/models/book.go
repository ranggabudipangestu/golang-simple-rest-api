package models

import (
	"github.com/ranggabudipangestu/go-book-store/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func (b *Book) GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func (b *Book) GetBookById(id int) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("id=?", id).Find(&getBook)
	return &getBook, db
}

func (b *Book) DeleteBook(id int) *Book {
	var book Book
	db.Where("id=?", id).Delete(&book)
	return &book
}
