package handlers

import (
	"sync"
	"toko/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// contoh lama
// func ListStokProducts(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var StokProducts []models.StokProduct
// 		db.Find(&StokProducts)
// 		c.JSON(200, StokProducts)
// 	}
// }

// contoh baru dengan goroutine

// @Summary Get a list of StokProducts
// @Description Retrieve a list of StokProducts from the database.
// @Produce json
// @Success 200 {array} StokProduct
// @Router /StokProducts [get]
func ListStok(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var StokProducts []models.StokProduct
		var wg sync.WaitGroup

		// Menambahkan satu goroutine ke WaitGroup
		wg.Add(1)

		// Memulai goroutine untuk melakukan operasi yang membutuhkan waktu lama
		go func() {
			defer wg.Done() // Menandai bahwa goroutine telah selesai
			db.Find(&StokProducts)
		}()

		// Menunggu goroutine selesai
		wg.Wait()

		c.JSON(200, StokProducts)
	}
}

func GetStokProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var StokProduct models.StokProduct
		if err := db.First(&StokProduct, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Stok not found"})
			return
		}
		c.JSON(200, StokProduct)
	}
}

func CreateStokProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.StokProduct
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		db.Create(&input)
		c.JSON(201, input)
	}
}

func UpdateStokProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var StokProduct models.StokProduct
		if err := db.First(&StokProduct, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "StokProduct not found"})
			return
		}

		var input models.StokProduct
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		db.Model(&StokProduct).Updates(input)
		c.JSON(200, StokProduct)
	}
}

func DeleteStokProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var StokProduct models.StokProduct
		if err := db.First(&StokProduct, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "StokProduct not found"})
			return
		}

		db.Delete(&StokProduct)
		c.JSON(200, gin.H{"message": "StokProduct deleted"})
	}
}
