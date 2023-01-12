package controller

import (
	"encoding/json"
	"github.com/bookSystem/model"
	"github.com/bookSystem/service"
	"github.com/bookSystem/util"
	"net/http"
)

type OrderController interface {
	CreateOrder(w http.ResponseWriter, r *http.Request)
	GetOrders(w http.ResponseWriter, r *http.Request)
}

type orderController struct {
	service service.OrderService
}

func (c orderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order model.Order
	json.NewDecoder(r.Body).Decode(&order)
	Order, err := c.service.CreateOrder(&order)
	if err != nil {
		util.ErrorJSON(w, err)
	} else {
		util.WriteJSON(w, 200, Order, "order")
	}
}

func (c orderController) GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	categories, err := c.service.GetOrders()
	if err != nil {
		util.ErrorJSON(w, err)
	} else {
		util.WriteJSON(w, 200, categories, "categories")
	}
}

func NewOrderController(orderService service.OrderService) OrderController {
	return &orderController{orderService}
}
