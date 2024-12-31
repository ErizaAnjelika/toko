package models

type Transaction struct {
	ID               uint    `gorm:"primary_key" json:"id"`
	UserID           uint    `json:"user_id"`
	Amount           float64 `json:"amount"`
	TransactionDate  string  `json:"transaction_date"`
	MetodePembayaran string  `json:"metode_pembayaran"`

	Items []TransactionItem `gorm:"foreignkey:TransactionID" json:"items"`
	User  User              `gorm:"foreignkey:UserID;references:ID" json:"user"`
}
