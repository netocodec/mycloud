package auth

import (
	"time"

	"../../db"
	"github.com/dgrijalva/jwt-go"
)

type membershipToken struct {
	User db.UsersList
	jwt.StandardClaims
}

var signKey []byte = []byte("MyCLoud#2020#sErver")

func GenerateToken(user db.UsersList) string {
	tokenResult := "TOKEN_ERR"
	tokenClaim := membershipToken{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    "MyCloud",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaim)
	if tokenStr, tokenStrErr := token.SignedString(signKey); tokenStrErr == nil {
		tokenResult = tokenStr
	}

	return tokenResult
}

func DecodeToken(tokenString string) (bool, db.UsersList) {
	var result db.UsersList
	var hasError = true

	token, err := jwt.ParseWithClaims(tokenString, &membershipToken{}, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})

	if err == nil {
		if claims, ok := token.Claims.(*membershipToken); ok && token.Valid {
			result = claims.User
			hasError = false
		}
	}

	return hasError, result
}
