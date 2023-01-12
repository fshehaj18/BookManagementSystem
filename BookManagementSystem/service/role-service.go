package service

import (
	"github.com/bookSystem/model"
	"github.com/bookSystem/repository"
)

type RoleService struct {
	repository repository.RoleRepository
}

func NewRoleService(repository repository.RoleRepository) *RoleService {
	return &RoleService{repository}
}

type RoleServiceInterface interface {
	CreateRole(id int) (*model.Role, error)
}

func (r *RoleService) GetRole(id int) (*model.Role, error) {
	role, err := r.repository.GetRole(id)
	if err != nil {
		return nil, err
	}
	return &role, nil
}
func (r *RoleService) CreateRole(role *model.Role) (*model.Role, error) {

	NewRole := model.Role{RoleName: role.RoleName}
	_, err := r.repository.AddRole(NewRole)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (u *RoleService) GetRoles() ([]model.Role, error) {
	roles, err := u.repository.GetRoles()
	if err != nil {
		return nil, err
	}
	return roles, nil
}
