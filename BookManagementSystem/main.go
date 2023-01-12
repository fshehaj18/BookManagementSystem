package main

import (
	"github.com/bookSystem/config"
	"github.com/bookSystem/controller"
	"github.com/bookSystem/repository"
	"github.com/bookSystem/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	db, err := config.NewConnection()
	if err != nil {
		log.Fatal("error")
	}

	r := mux.NewRouter()
	userRepository := repository.NewUserRepo(db)
	roleRepository := repository.NewRoleRepo(db)
	categoryRepository := repository.NewCategoryRepo(db)
	bookRepository := repository.NewBookRepo(db)
	orderRepository := repository.NewOrderRepo(db)

	userService := service.NewUserService(userRepository)
	roleService := service.NewRoleService(roleRepository)
	categoryService := service.NewCategoryService(categoryRepository)
	bookService := service.NewBookService(bookRepository)
	orderService := service.NewOrderService(orderRepository)

	userController := controller.NewController(*userService, *roleService)
	categoryController := controller.NewCategoryController(*categoryService)
	bookController := controller.NewBookController(*bookService)
	orderController := controller.NewOrderController(*orderService)

	r.HandleFunc("/users/", userController.CreateUser).Methods("POST")
	r.HandleFunc("/users/", userController.GetUsers).Methods("GET")
	r.HandleFunc("/roles/", userController.GetRoles).Methods("GET")
	r.HandleFunc("/roles/", userController.GetRoles).Methods("POST")
	r.HandleFunc("/categories/", categoryController.CreateCategory).Methods("POST")
	r.HandleFunc("/categories/", categoryController.GetCategories).Methods("GET")
	r.HandleFunc("/books/", bookController.CreateBook).Methods("POST")
	r.HandleFunc("/books/", bookController.GetBooks).Methods("GET")
	r.HandleFunc("/books/{categoryId}", bookController.GetBooksByCategoryId).Methods("GET")
	r.HandleFunc("/orders/", orderController.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/", orderController.GetOrders).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", r))
}
