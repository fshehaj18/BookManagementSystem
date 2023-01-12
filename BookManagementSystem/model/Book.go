package model

import (
	"time"
)

type Book struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	PublishDate time.Time  `json:"publish_date"`
	Author      string     `json:"author"`
	Image       string     `json:"image"`
	Publisher   string     `json:"publisher"`
	Status      bool       `json:"status"`
	Categories  []Category `gorm:"many2many:book_category;association_jointable_foreignkey:category_id;jointable_foreignkey:book_id"`
}

func (Book) TableName() string {
	return "books"
}

type Order struct {
	OrderId    int       `json:"order_id"`
	OrderDate  time.Time `json:"order_date"`
	ReturnDate time.Time `json:"return_date"`
	BookId     int       `json:"book_id"`
	UserId     int       `json:"user_id"`
	User       User      `references: Id`
}

func (Order) TableName() string {
	return "order"
}
