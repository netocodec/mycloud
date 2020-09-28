package auth

import (
	"log"
	"testing"

	"../../db"
)

var user db.UsersList = db.UsersList{
	UserName:     "user_test_001",
	UserPassword: "user_pass_01234566789",
	UserID:       0,
	IsAdmin:      0,
}

func TestGeneratedToken(test *testing.T) {
	if tokenStr := GenerateToken(user); tokenStr == "TOKEN_ERR" {
		test.Errorf("Cannot generate token for the user: %s", user.UserName)
	}
}

func TestDecodedToken(test *testing.T) {
	if tokenStr := GenerateToken(user); tokenStr != "TOKEN_ERR" {
		if hasError, decToken := DecodeToken(tokenStr); !hasError {
			log.Printf("User token decoded with sucess!")
			log.Println(decToken)
		} else {
			test.Errorf("Cannot decode this token! (Token STR: %s)", tokenStr)
		}
	} else {
		test.Errorf("Cannot generate token for the user: %s", user.UserName)
	}
}
