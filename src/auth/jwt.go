package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret = os.Getenv("JWT_SECRET")
var jwtKey = []byte(JWTSecret)

type Claims struct {
	UserID int
	jwt.RegisteredClaims
}

func GenerateJWT(user_id int) (string, error) {
	claims := Claims{
		UserID: user_id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2400 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (any, error) { return jwtKey, nil })
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
