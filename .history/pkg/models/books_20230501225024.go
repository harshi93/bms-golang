package models

import (
	"github.com/harshi93/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"json:"name""`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// we are receiving an object of type Book
// we returning an object of type book
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// slice is being used because we want to return slice
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

//takes in id for book
/* we are returning db object that is of type *gorm.DB and
we are returning getBook which is of type Book
*/
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book //var getbook is of type Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}

//routes -> book controllers -> books.go
