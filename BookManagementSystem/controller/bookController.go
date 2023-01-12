package controller

import (
	"encoding/json"
	"fmt"
	"github.com/bookSystem/model"
	"github.com/bookSystem/service"
	"github.com/bookSystem/util"
	"net/http"
)

type BookController interface {
	CreateBook(w http.ResponseWriter, r *http.Request)
	GetBooks(w http.ResponseWriter, r *http.Request)
	GetBooksByCategoryId(w http.ResponseWriter, r *http.Request)
}

type bookController struct {
	service service.BookService
}

func (c bookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book model.Book
	json.NewDecoder(r.Body).Decode(&book)
	_, err := c.service.CreateBook(&book)
	if err != nil {
		util.ErrorJSON(w, err)
	} else {
		util.WriteJSON(w, 200, book, "book")
	}
}

func (c bookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books, err := c.service.GetBooks()
	if err != nil {
		util.ErrorJSON(w, err)
	} else {
		util.WriteJSON(w, 200, books, "books")
	}
}
func (c bookController) GetBooksByCategoryId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("%+v", r.URL.Query().Get("categoryId"))
	books, err := c.service.GetBooksByCategoryId(4)
	if err != nil {
		util.ErrorJSON(w, err)
	} else {
		util.WriteJSON(w, 200, books, "books")
	}
}
func NewBookController(bookService service.BookService) BookController {
	return &bookController{bookService}
}
