package handlers

import "github.com/dgrijalva/jwt-go"

var jwtKey = []byte("your-secret-key") // Ganti dengan kunci rahasia Anda

func CreateToken(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["role"] = role // Tambahkan klaim role
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
