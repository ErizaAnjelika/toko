package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	ID                uint    `gorm:"primary_key" json:"id"`
	Name              string  `json:"name"`
	CategoryID        uint    `json:"category_id"`
	HargaJual         float64 `json:"harga_jual"`
	HargaBeli         float64 `json:"harga_beli"`
	Quantity          uint    `json:"quantity"`
	TanggalMasuk      string  `json:"tanggal_masuk"`
	TanggalKadaluarsa string  `json:"tanggal_kadaluarsa"`
	MarginKeuntungan  float64 `json:"margin_keuntungan"`
	Barcode           string  `gorm:"unique;not null" json:"barcode"`

	Category ProductCategory `gorm:"foreignkey:CategoryID;references:ID" json:"category"`
}

// BeforeSave GORM hook untuk melakukan logika sebelum data disimpan
func (p *Product) BeforeSave(tx *gorm.DB) (err error) {
	// Jika HargaJual belum dihitung, hitung otomatis
	if p.HargaJual == 0 && p.HargaBeli > 0 && p.MarginKeuntungan > 0 {
		p.HargaJual = p.HargaBeli + (p.HargaBeli * p.MarginKeuntungan / 100)
	} else {
		p.HargaJual = p.HargaBeli + (p.HargaBeli * p.MarginKeuntungan / 100)
	}

	return nil
}
