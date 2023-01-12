package model

type User struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	RoleId    uint   `json:"role_id"`
	Role      Role   `gorm:"references:RoleId"`
}

func (user *User) TableName() string {
	return "user"
}

type Role struct {
	RoleId   uint   `json:"role_id" gorm:"primaryKey"`
	RoleName string `json:"role_name"`
}

func (role *Role) TableName() string {
	return "role"
}
