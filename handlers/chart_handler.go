package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetSalesByDate(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		type SalesData struct {
			Tanggal        string  `json:"tanggal"`
			TotalPenjualan float64 `json:"total_penjualan"`
		}

		var sales []SalesData
		if err := db.Raw(`
			SELECT 
				DATE(transaction_date) AS tanggal, 
				SUM(amount) AS total_penjualan
			FROM transactions
			GROUP BY DATE(transaction_date)
			ORDER BY tanggal ASC
		`).Scan(&sales).Error; err != nil {
			c.JSON(500, gin.H{"message": "Error fetching sales data"})
			return
		}

		c.JSON(200, gin.H{"data": sales})
	}
}

func GetTopSellingProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		type TopProduct struct {
			NamaProduk    string `json:"nama_produk"`
			JumlahTerjual int    `json:"jumlah_terjual"`
		}

		var products []TopProduct
		if err := db.Raw(`
			SELECT 
				p.name AS nama_produk, 
				SUM(ti.quantity) AS jumlah_terjual
			FROM transaction_items ti
			JOIN products p ON ti.product_id = p.id
			GROUP BY p.id
			ORDER BY jumlah_terjual DESC
			LIMIT 10
		`).Scan(&products).Error; err != nil {
			c.JSON(500, gin.H{"message": "Error fetching top selling products"})
			return
		}

		c.JSON(200, gin.H{"data": products})
	}
}
