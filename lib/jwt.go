package lib

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateduserIdToken(id int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"iat": jwt.NumericDate{
			Time: time.Now(),
		},
	})

	tokenSignedString, _ := token.SignedString([]byte("secret"))
	return tokenSignedString
}

func ValidateToken(token string) (bool, int) {
	validated, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodECDSA); ok {
			return nil, fmt.Errorf("unnexpected signing method: %v", t.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if err != nil {
		panic("Error: Token Invalid")
	}
	if claims, ok := validated.Claims.(jwt.MapClaims); ok {
		userId := int(claims["id"].(float64))
		return true, userId
	} else {
		panic("error: token invalid")
	}
}
