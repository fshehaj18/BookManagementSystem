package service

import (
	"github.com/bookSystem/model"
	"github.com/bookSystem/repository"
)

type UserServiceInterface interface {
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(int, user *model.User) (*model.User, error)
	GetUsers() ([]model.User, error)
	GetRoles() ([]model.Role, error)
}

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
	return &UserService{repository}
}

func (u *UserService) CreateUser(user *model.User) (*model.User, error) {

	NewUser := model.User{FirstName: user.FirstName, LastName: user.LastName, Email: user.Email, Phone: user.Phone, RoleId: user.RoleId}
	_, err := u.repository.AddUser(NewUser)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) Update(id int, user *model.User) (*model.User, error) {

	NewUser := model.User{FirstName: user.FirstName, LastName: user.LastName, Email: user.Email, Phone: user.Phone, RoleId: user.RoleId}
	_, err := u.repository.UpdateUser(id, NewUser)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) GetUsers() ([]model.User, error) {
	users, err := u.repository.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
