package service

import (
	"github.com/bookSystem/model"
	"github.com/bookSystem/repository"
)

type CategoryService struct {
	repository repository.CategoryRepository
}

func NewCategoryService(repository repository.CategoryRepository) *CategoryService {
	return &CategoryService{repository}
}

func (c *CategoryService) CreateCategory(category *model.Category) (*model.Category, error) {
	NewCategory := model.Category{CategoryName: category.CategoryName}
	_, err := c.repository.AddCategory(NewCategory)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryService) GetCategories() ([]model.Category, error) {
	categories, err := c.repository.GetCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}
