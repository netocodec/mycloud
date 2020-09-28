package db

import (
	"fmt"
	"log"
	"testing"
)

const userTestName string = "usertest001"
const userTestPass string = "userpass0123456789"

func TestDBInit(test *testing.T) {
	InsertUser(userTestName, userTestPass, 0)

	if users := GetAllUsers(); len(users) == 0 {
		test.Error("There is no user to show!")
	} else {
		user := GetUserByName(userTestName)

		log.Printf("Testing user %s, delete operation!", user.UserName)
		if user.UserName == userTestName {
			if isUserDeleted := DeleteUserByID(user.UserID); !isUserDeleted {
				test.Errorf("The user %s cannot be deleted! (Reason: Maybe it does not exists in DB)", user.UserName)
			}
		} else {
			test.Errorf("Cannot get the user %s", user.UserName)
		}
	}
}

func TestUserLogin(test *testing.T) {
	InsertUser(userTestName, userTestPass, 0)

	log.Printf("User %s is try to login...", userTestName)
	if isLoginSucess, userMembership := LoginMembership(userTestName, userTestPass); !isLoginSucess {
		test.Errorf("The user %s does not have the correct credencials!", userTestName)
	} else {
		log.Printf("User %s login success! (Username: %s | IsAdmin: %d)", userTestName, userMembership.UserName, userMembership.IsAdmin)
	}
}

func TestUserFailureLogin(test *testing.T) {
	InsertUser(userTestName, userTestPass, 0)

	log.Printf("User %s is try to login...", userTestName)
	if isLoginSucess, _ := LoginMembership(userTestName, fmt.Sprintf("%s123", userTestPass)); isLoginSucess {
		test.Errorf("The user %s have the correct credencials, it is suposse not to have it!", userTestName)
	} else {
		log.Printf("User %s login fail correct!", userTestName)
	}
}

func BenchmarkDBInsert(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		InsertUser(fmt.Sprintf("USER_%d_BENCH", i), "userpass", 0)
	}
}
