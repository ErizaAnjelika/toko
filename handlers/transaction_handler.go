package handlers

import (
	"time"
	"toko/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ListTransactions(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var transaction []models.Transaction

		// Memuat produk beserta kategori terkait menggunakan Preload
		if err := db.Preload("Items.Product.Category").Find(&transaction).Error; err != nil {
			c.JSON(500, gin.H{"message": "Error retrieving transaction", "error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": transaction})
	}
}

func GetTransactionWithItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var transaction models.Transaction
		if err := db.Preload("Items.Product.Category").First(&transaction, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Transaction not found"})
			return
		}

		c.JSON(200, gin.H{"data": transaction})
	}
}

// func CreateTransaction(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var input models.Transaction
// 		if err := c.ShouldBindJSON(&input); err != nil {
// 			c.JSON(400, gin.H{"message": "Invalid input"})
// 			return
// 		}

// 		// Jika Anda ingin mengaitkan item transaksi, pastikan item-item tersebut ada dalam database
// 		for _, item := range input.Items {
// 			var product models.Product
// 			if err := db.First(&product, item.ProductID).Error; err != nil {
// 				c.JSON(400, gin.H{"message": "Invalid product ID"})
// 				return
// 			}
// 		}

// 		db.Create(&input)
// 		c.JSON(201, input)
// 	}
// }

func CreateTransaction(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		type ItemInput struct {
			Barcode  string `json:"barcode"`
			Quantity uint   `json:"quantity"`
		}

		type TransactionInput struct {
			UserID           uint        `json:"user_id"`
			MetodePembayaran string      `json:"metode_pembayaran"`
			Items            []ItemInput `json:"items"`
		}

		var input TransactionInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		var transaction models.Transaction
		transaction.UserID = input.UserID
		transaction.MetodePembayaran = input.MetodePembayaran
		transaction.TransactionDate = time.Now().Format("2006-01-02")

		totalAmount := 0.0
		var items []models.TransactionItem

		for _, itemInput := range input.Items {
			var product models.Product
			if err := db.Where("barcode = ?", itemInput.Barcode).First(&product).Error; err != nil {
				c.JSON(400, gin.H{"message": "Invalid product barcode"})
				return
			}

			if product.Quantity < itemInput.Quantity {
				c.JSON(400, gin.H{"message": "Insufficient product stock", "product": product.Name})
				return
			}

			// Update stock produk
			product.Quantity -= itemInput.Quantity
			if err := db.Save(&product).Error; err != nil {
				c.JSON(500, gin.H{"message": "Failed to update product stock"})
				return
			}

			itemAmount := float64(itemInput.Quantity) * product.HargaJual
			totalAmount += itemAmount

			items = append(items, models.TransactionItem{
				ProductID: product.ID,
				Quantity:  itemInput.Quantity,
				Price:     product.HargaJual,
				Amount:    itemAmount,
			})
		}

		transaction.Amount = totalAmount
		transaction.Items = items

		// Simpan transaksi dan item terkait
		if err := db.Create(&transaction).Error; err != nil {
			c.JSON(500, gin.H{"message": "Failed to create transaction"})
			return
		}

		c.JSON(201, gin.H{"message": "Transaction created successfully", "data": transaction})
	}
}
