package repository

import (
	"fmt"
	"github.com/bookSystem/model"
	"gorm.io/gorm"
)

type orderRepo struct {
	DB *gorm.DB
}

type OrderRepository interface {
	GetOrder(id int) (model.Order, error)
	GetOrders() ([]model.Order, error)
	AddOrder(newOrder model.Order) (model.Order, error)
}

func (r *orderRepo) GetOrder(orderId int) (model.Order, error) {
	var order model.Order
	err := r.DB.First(orderId).Error
	if err != nil {
		order = model.Order{}
		return order, err
	}
	return order, nil
}
func (r *orderRepo) AddOrder(newOrder model.Order) (model.Order, error) {
	err := r.DB.Create(&newOrder).Error
	if err != nil {
		newOrder = model.Order{}
		return newOrder, err
	}
	return newOrder, nil
}
func (r *orderRepo) GetOrders() ([]model.Order, error) {
	var orders []model.Order
	r.DB.Find(&orders)
	fmt.Printf("%+v\n", orders)
	return orders, nil
}
func NewOrderRepo(db *gorm.DB) OrderRepository {
	repo := new(orderRepo)
	repo.DB = db
	return repo
}
