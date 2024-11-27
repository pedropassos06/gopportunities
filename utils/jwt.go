package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
)

func GenerateJWT(userInfo map[string]interface{}) (string, error) {
	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userInfo["id"],
		"email":   userInfo["email"],
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	// sign the token with jwt secret
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validando o algoritmo de assinatura
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return false
	}

	// Verifica se o token é válido e se as claims são confiáveis
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp := claims["exp"].(float64)
		return time.Unix(int64(exp), 0).After(time.Now()) // Verifica expiração
	}
	return false
}
