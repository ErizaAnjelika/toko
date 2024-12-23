package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	ID         uint    `gorm:"primary_key" json:"id"`
	Name       string  `json:"name"`
	CategoryID uint    `json:"category_id"`
	Price      float64 `json:"price"` // Tambahkan atribut Price di sini
}