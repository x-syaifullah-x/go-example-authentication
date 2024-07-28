package jwt

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/x-syaifullah-x/go-crud/src/pkg/logger"
)

// type claims struct {
// 	jwt.RegisteredClaims
// }

var jwtSecret = []byte("1NGIzSxvuWnYmEr7+hauVsG1fRln5hD+PqI7MBLavU4=")

func GenerateJWT(id string) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
		ID:        id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateJWT(token string) (*jwt.RegisteredClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := jwtToken.Claims.(*jwt.RegisteredClaims); ok && jwtToken.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func generateSecretKey() string {
	key := make([]byte, 32) // 256 bits
	if _, err := rand.Read(key); err != nil {
		logger.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(key)
}
