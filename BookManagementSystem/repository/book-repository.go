package repository

import (
	"fmt"
	"github.com/bookSystem/model"
	"gorm.io/gorm"
)

type bookRepo struct {
	DB *gorm.DB
}

type BookRepository interface {
	AddBook(book model.Book) (model.Book, error)
	FindBookById(id int) (model.Book, error)
	GetBooks() ([]model.Book, error)
	GetBooksByCategory(categoryId int) ([]model.Book, error)
}

func (b *bookRepo) AddBook(newBook model.Book) (model.Book, error) {
	b.DB.Model(&newBook).Association("Categories")
	err := b.DB.Omit("Categories").Create(&newBook).Debug().Error
	if err != nil {
		newBook = model.Book{}
		return newBook, err
	}
	return newBook, nil
}
func (b *bookRepo) FindBookById(id int) (model.Book, error) {
	var book model.Book
	err := b.DB.First(id).Error
	if err != nil {
		book = model.Book{}
		return book, err
	}
	return book, nil
}
func (b *bookRepo) GetBooks() ([]model.Book, error) {
	var books []model.Book
	b.DB.Table("books").Debug().Find(&books)
	return books, nil
}
func (b *bookRepo) GetBooksByCategory(categoryId int) ([]model.Book, error) {
	var books []model.Book
	b.DB.Table("books").Joins("left join book_category on books.id=book_category.book_id").Where("book_category.category_id", categoryId).Find(&books)
	fmt.Printf("%+v ekcnwp", books)
	return books, nil
}
func NewBookRepo(db *gorm.DB) BookRepository {
	repo := new(bookRepo)
	repo.DB = db
	return repo
}
