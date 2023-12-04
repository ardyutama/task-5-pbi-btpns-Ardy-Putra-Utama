package helpers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var signingKey = []byte("tes123")

func GenerateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(signingKey)
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
}

func ExtractUserID(c *gin.Context) uint {
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		return 0
	}

	userID, ok := userIDInterface.(uint)
	if !ok {
		return 0
	}

	return userID
}
