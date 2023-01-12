package service

import (
	"github.com/bookSystem/model"
	"github.com/bookSystem/repository"
)

type BookService struct {
	repository repository.BookRepository
}

func NewBookService(repository repository.BookRepository) *BookService {
	return &BookService{repository}
}

type BookServiceInterface interface {
	CreateBook(user *model.Book) (*model.Book, error)
	GetBooks() ([]model.Book, error)
	GetBooksByCategoryId() ([]model.Book, error)
}

func (b *BookService) CreateBook(Book *model.Book) (*model.Book, error) {
	NewBook := model.Book{Description: Book.Description, Title: Book.Title, Price: Book.Price,
		Author: Book.Author, PublishDate: Book.PublishDate, Status: false, Publisher: Book.Publisher, Image: Book.Image,
		Categories: Book.Categories,
	}
	_, err := b.repository.AddBook(NewBook)
	if err != nil {
		return nil, err
	}
	return Book, nil
}
func (b *BookService) GetBooksByCategoryId(categoryId int) ([]model.Book, error) {
	books, err := b.repository.GetBooksByCategory(categoryId)
	if err != nil {
		return nil, err
	}
	return books, nil
}
func (b *BookService) GetBooks() ([]model.Book, error) {
	books, err := b.repository.GetBooks()
	if err != nil {
		return nil, err
	}
	return books, nil
}
