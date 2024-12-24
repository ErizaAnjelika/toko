package models

type ProductCategory struct{
	ID uint `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}