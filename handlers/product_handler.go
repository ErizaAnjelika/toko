package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
	"toko/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// contoh lama
// func ListProducts(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var products []models.Product
// 		db.Find(&products)
// 		c.JSON(200, products)
// 	}
// }

// contoh baru dengan goroutine

// @Summary Get a list of products
// @Description Retrieve a list of products from the database.
// @Produce json
// @Success 200 {array} Product
// @Router /products [get]

// GenerateBarcode generates a unique barcode based on product details
func GenerateBarcode(name string, categoryID uint, tanggalMasuk string) string {
	data := fmt.Sprintf("%s-%d-%s-%d", name, categoryID, tanggalMasuk, time.Now().UnixNano())
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])[:12] // Ambil 12 karakter pertama untuk barcode
}

func ListProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []models.Product

		// Memuat produk beserta kategori terkait menggunakan Preload
		if err := db.Preload("Category").Find(&products).Error; err != nil {
			c.JSON(500, gin.H{"message": "Error retrieving products", "error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": products})
	}
}

func GetProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product models.Product
		if err := db.Preload("Category").First(&product, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Product not found"})
			return
		}
		c.JSON(200, gin.H{
			"data": product,
		})
	}
}

func CreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Product

		// Validasi input JSON
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input", "error": err.Error()})
			return
		}

		// Generate barcode jika belum ada
		if input.Barcode == "" {
			input.Barcode = GenerateBarcode(input.Name, input.CategoryID, input.TanggalMasuk)
		}

		// Simpan data ke database
		if err := db.Create(&input).Error; err != nil {
			c.JSON(500, gin.H{"message": "Failed to create product", "error": err.Error()})
			return
		}

		c.JSON(201, gin.H{"message": "Product created successfully", "data": input})
	}
}

func UpdateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product models.Product
		if err := db.First(&product, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Product not found"})
			return
		}

		var input models.Product
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		db.Model(&product).Updates(input)
		c.JSON(200, gin.H{"message": "Product updated successfully", "data": product})
	}
}

func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product models.Product
		if err := db.First(&product, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Product not found"})
			return
		}

		db.Delete(&product)
		c.JSON(200, gin.H{"message": "Product deleted"})
	}
}
