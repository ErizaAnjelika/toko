package handlers

import (
	"toko/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		var user models.User
		// Cek apakah username ada di database
		if err := db.Where("username = ?", input.Username).First(&user).Error; err != nil {
			c.JSON(401, gin.H{"message": "Invalid username or password"})
			return
		}

		// Verifikasi password
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			c.JSON(401, gin.H{"message": "Invalid username or password"})
			return
		}

		// Pastikan untuk memeriksa kata sandi yang benar di sini
		token, err := CreateToken(user.ID, user.Role)
		if err != nil {
			c.JSON(500, gin.H{"message": "Internal Server Error"})
			return
		}

		c.JSON(200, gin.H{
			"message": "Login successful",
			"token":   token,
			"data": map[string]interface{}{
				"id":         user.ID,
				"username":   user.Username,
				"email":      user.Email,
				"nama_kasir": user.NamaKasir,
				"role":       user.Role,
			},
		})
	}
}
