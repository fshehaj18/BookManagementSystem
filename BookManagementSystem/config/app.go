package config

import (
	"github.com/bookSystem/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:1234@tcp(localhost:3306)/trip-management-system"

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func NewConnection() (*gorm.DB, error) {
	var dsn = "root:1234@tcp(localhost:3306)/book-management-system"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	err = db.AutoMigrate(&model.User{}, &model.Role{})
	if err != nil {
		return nil, err
	}
	if err != nil {
		return db, err
	}
	return db, nil
}
