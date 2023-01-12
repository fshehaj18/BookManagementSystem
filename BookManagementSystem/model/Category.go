package model

type Category struct {
	Id           uint   `json:"id"`
	CategoryName string `json:"category_name"`
}

func (category *Category) TableName() string {
	return "categories"
}
