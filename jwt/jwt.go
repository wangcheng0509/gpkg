package jwt

import (
	"encoding/json"
	"errors"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte

type Claims struct {
	Unique_name  string `json:"unique_name"`
	Guid         string `json:"guid"`
	Avatar       string `json:"avatar"`
	DisplayName  string `json:"displayName"`
	LoginName    string `json:"loginName"`
	EmailAddress string `json:"emailAddress"`
	UserType     string `json:"userType"`
	Time         string `json:"time"`
	jwt.StandardClaims
}

func Setup(_jwtSecret string) {
	jwtSecret = []byte(_jwtSecret)
}

// 以内置对象生成token
func GenerateTokenByBuiltin(claims Claims) (string, error) {
	return GenerateToken(claims)
}

// 自定义对象生成token
func GenerateToken(claims interface{}) (string, error) {
	if tmp, ok := claims.(jwt.Claims); ok {
		tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, tmp)
		token, err := tokenClaims.SignedString(jwtSecret)
		return token, err
	}
	return "", errors.New("claims is not jwt.Claims")
}

// 解析成内置对象
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	fmt.Println(tokenClaims)
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// 解析成json字符串
func Parse(token string) (string, error) {
	tokenRst, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenRst != nil {
		if tokenRst.Valid {
			claimsJson, _ := json.Marshal(tokenRst.Claims)
			return string(claimsJson), nil
		}
	}
	return "", err
}
