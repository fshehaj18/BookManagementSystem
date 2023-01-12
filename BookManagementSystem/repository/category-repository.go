package repository

import (
	"github.com/bookSystem/model"
	"gorm.io/gorm"
)

type categoryRepo struct {
	DB *gorm.DB
}

type CategoryRepository interface {
	AddCategory(category model.Category) (model.Category, error)
	FindCategoryById(id int) (model.Category, error)
	GetCategories() ([]model.Category, error)
}

func (b *categoryRepo) AddCategory(newCategory model.Category) (model.Category, error) {
	err := b.DB.Create(&newCategory).Error
	if err != nil {
		newCategory = model.Category{}
		return newCategory, err
	}
	return newCategory, nil
}

func (b *categoryRepo) FindCategoryById(id int) (model.Category, error) {
	var category model.Category
	err := b.DB.First(id).Error
	if err != nil {
		category = model.Category{}
		return category, err
	}
	return category, nil
}

func (b *categoryRepo) GetCategories() ([]model.Category, error) {
	var categories []model.Category
	b.DB.Table("categories").Find(&categories)
	return categories, nil
}

func NewCategoryRepo(db *gorm.DB) CategoryRepository {
	repo := new(categoryRepo)
	repo.DB = db
	return repo
}
