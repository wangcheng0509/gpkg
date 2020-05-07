package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte
var issuer string

type Claims struct {
	InfoJson string
	jwt.StandardClaims
}

func Setup(_jwtSecret, _issuer string) {
	jwtSecret = []byte(_jwtSecret)
	issuer = _issuer
}

// GenerateToken generate tokens used for auth
func GenerateToken(infoJson string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		infoJson,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
