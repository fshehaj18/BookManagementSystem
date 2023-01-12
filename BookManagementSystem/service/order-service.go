package service

import (
	"github.com/bookSystem/model"
	"github.com/bookSystem/repository"
)

type OrderService struct {
	repository     repository.OrderRepository
	userRepository repository.UserRepository
}

func NewOrderService(repository repository.OrderRepository, userRepository repository.UserRepository) *OrderService {
	return &OrderService{repository, userRepository}
}

type OrderServiceInterface interface {
	CreateOrder(order *model.Order) (*model.Order, error)
	GetOrders() ([]model.Order, error)
	GetOrdersByCategoryId() ([]model.Order, error)
}

func (b *OrderService) CreateOrder(Order *model.Order) (*model.Order, error) {
	NewOrder := model.Order{OrderDate: Order.OrderDate,
		ReturnDate: Order.ReturnDate, UserId: Order.UserId, BookId: Order.BookId}
	_, err := b.repository.AddOrder(NewOrder)
	b.userRepository.UpdateUser(Order.UserId)
	if err != nil {
		return nil, err
	}
	return Order, nil
}

func (b *OrderService) GetOrders() ([]model.Order, error) {
	orders, err := b.repository.GetOrders()
	if err != nil {
		return nil, err
	}
	return orders, nil
}
