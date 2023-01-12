package controller

import (
	"encoding/json"
	"github.com/bookSystem/model"
	"github.com/bookSystem/service"
	"github.com/bookSystem/util"
	"net/http"
)

type CategoryController interface {
	CreateCategory(w http.ResponseWriter, r *http.Request)
	GetCategories(w http.ResponseWriter, r *http.Request)
}

type categoryController struct {
	service service.CategoryService
}

func (c categoryController) CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var category model.Category
	json.NewDecoder(r.Body).Decode(&category)
	Category, err := c.service.CreateCategory(&category)
	if err != nil {
		util.ErrorJSON(w, err)
	} else {
		util.WriteJSON(w, 200, Category, "category")
	}
}

func (c categoryController) GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	categories, err := c.service.GetCategories()
	if err != nil {
		util.ErrorJSON(w, err)
	} else {
		util.WriteJSON(w, 200, categories, "categories")
	}
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &categoryController{categoryService}
}
