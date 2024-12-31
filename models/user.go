package models

type User struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	NamaKasir string `json:"nama_kasir"`
	Role      string `json:"role"` // e.g., "admin", "kasir"
}
