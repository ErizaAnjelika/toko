package middlewares

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("your-secret-key") // Ganti dengan kunci rahasia Anda

func AuthMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil token dari header Authorization
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		// Hapus prefix "Bearer " jika ada
		if strings.HasPrefix(tokenString, "Bearer ") {
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		}

		// Parsing token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		// Ambil klaim dari token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		// Ambil user_id dari klaim token
		userID, ok := claims["user_id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token payload"})
			c.Abort()
			return
		}
		c.Set("user_id", uint(userID))

		// Ambil role dari klaim token
		userRole, ok := claims["role"].(string)
		if !ok || userRole == "" {
			c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden: role not found in token"})
			c.Abort()
			return
		}

		// Periksa apakah role pengguna diizinkan
		if len(allowedRoles) > 0 {
			roleAllowed := false
			for _, role := range allowedRoles {
				if userRole == role {
					roleAllowed = true
					break
				}
			}

			if !roleAllowed {
				c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden: role not allowed"})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
