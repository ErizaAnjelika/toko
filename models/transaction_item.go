package models

type TransactionItem struct {
	ID            uint    `gorm:"primary_key" json:"id"`
	TransactionID uint    `json:"transaction_id"`
	ProductID     uint    `json:"product_id"`
	Quantity      uint    `json:"quantity"`
	Price         float64 `json:"price"`
	Amount        float64 `json:"amount"`

	Product Product `gorm:"foreignkey:ProductID;references:ID" json:"product"`
}
