package repository

import (
	"fmt"
	"github.com/bookSystem/model"
	"gorm.io/gorm"
)

type userRepo struct {
	DB *gorm.DB
}

type UserRepository interface {
	AddUser(model.User) (model.User, error)
	UpdateUser(int, model.User) (model.User, error)
	FindUserById(id int) (model.User, error)
	GetUsers() ([]model.User, error)
}

func (u *userRepo) AddUser(newUser model.User) (model.User, error) {
	err := u.DB.Create(&newUser).Error
	if err != nil {
		newUser = model.User{}
		return newUser, err
	}
	return newUser, nil
}
func (u *userRepo) UpdateUser(id int, newUser model.User) (model.User, error) {
	newUser.Id = uint(id)
	err := u.DB.Create(&newUser).Error
	if err != nil {
		newUser = model.User{}
		return newUser, err
	}
	return newUser, nil
}
func (u *userRepo) FindUserById(id int) (model.User, error) {
	var user model.User
	err := u.DB.First(&user, id).Error
	if err != nil {
		user = model.User{}
		return user, err
	}
	return user, nil
}
func (u *userRepo) GetUsers() ([]model.User, error) {
	var users []model.User
	u.DB.Table("User").Preload("Role").Find(&users)
	fmt.Printf("%+v\n", users)
	return users, nil
}
func NewUserRepo(db *gorm.DB) UserRepository {
	repo := new(userRepo)
	repo.DB = db
	return repo
}
