package repository

import (
	"fmt"
	"github.com/bookSystem/model"
	"gorm.io/gorm"
)

type roleRepo struct {
	DB *gorm.DB
}

type RoleRepository interface {
	GetRole(id int) (model.Role, error)
	GetRoles() ([]model.Role, error)
	AddRole(newRole model.Role) (model.Role, error)
}

func (r *roleRepo) GetRole(roleId int) (model.Role, error) {
	var role model.Role
	err := r.DB.First(roleId).Error
	if err != nil {
		role = model.Role{}
		return role, err
	}
	return role, nil
}
func (r *roleRepo) AddRole(newRole model.Role) (model.Role, error) {
	err := r.DB.Create(&newRole).Error
	if err != nil {
		newRole = model.Role{}
		return newRole, err
	}
	return newRole, nil
}
func (r *roleRepo) GetRoles() ([]model.Role, error) {
	var roles []model.Role
	r.DB.Find(&roles)
	fmt.Printf("%+v\n", roles)
	return roles, nil
}
func NewRoleRepo(db *gorm.DB) RoleRepository {
	repo := new(roleRepo)
	repo.DB = db
	return repo
}
