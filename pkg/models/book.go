package models

import (
	"github.com/bh1995/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()

	println("book init() called")

	db.AutoMigrate(&Book{})
}

func (newBook *Book) CreateBook() *Book {
	db.NewRecord(newBook)

	db.Create(&newBook)

	return newBook
}

func GetAllBooks() []Book {
	var BookList []Book

	db.Find(&BookList)

	return BookList
}

func GetBookById(bookId int64) (*Book, *gorm.DB) {
	var aBook Book
	db := db.Where("ID=?", bookId).Find(&aBook)
	return &aBook, db
}

func DeleteBook(bookID int64) Book {
	var aBookToDelete Book

	db.Where("ID=?", bookID).Delete(aBookToDelete)

	return aBookToDelete
}
