package auth

import (
	"time"

	"../../db"
	"github.com/dgrijalva/jwt-go"
)

type MembershipToken struct {
	User db.UsersList
	jwt.StandardClaims
}

var SignKey []byte = []byte("MyCLoud#2020#sErver")

func GenerateToken(user db.UsersList) string {
	tokenResult := "TOKEN_ERR"
	tokenClaim := MembershipToken{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    "MyCloud",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaim)
	if tokenStr, tokenStrErr := token.SignedString(SignKey); tokenStrErr == nil {
		tokenResult = tokenStr
	}

	return tokenResult
}

func DecodeToken(tokenString string) (bool, db.UsersList) {
	var result db.UsersList
	var hasError = true

	token, err := jwt.ParseWithClaims(tokenString, &MembershipToken{}, func(token *jwt.Token) (interface{}, error) {
		return SignKey, nil
	})

	if err == nil {
		if claims, ok := token.Claims.(*MembershipToken); ok && token.Valid {
			result = claims.User
			hasError = false
		}
	}

	return hasError, result
}
