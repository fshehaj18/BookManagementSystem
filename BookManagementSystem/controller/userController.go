package controller

import (
	"encoding/json"
	"github.com/bookSystem/model"
	"github.com/bookSystem/service"
	"github.com/bookSystem/util"
	"net/http"
)

type Controller interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetRoles(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	service     service.UserService
	roleService service.RoleService
}

func (c controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	User, err := c.service.CreateUser(&user)
	if err != nil {
		util.ErrorJSON(w, err)
	} else {
		util.WriteJSON(w, 200, User, "user")
	}
}

func (c controller) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := c.service.GetUsers()
	if err != nil {
		util.ErrorJSON(w, err)
	} else {
		util.WriteJSON(w, 200, users, "users")
	}
}

func (c controller) GetRoles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := c.roleService.GetRoles()
	if err != nil {
		util.ErrorJSON(w, err)
	} else {
		util.WriteJSON(w, 200, users, "roles")
	}
}

func NewController(userService service.UserService, roleService service.RoleService) Controller {
	return &controller{userService, roleService}
}
