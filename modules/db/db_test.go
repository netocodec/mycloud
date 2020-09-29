package db

import (
	"fmt"
	"log"
	"testing"
)

const userTestName string = "usertest001"
const userTestPass string = "userpass0123456789"
const editUserTestPass string = "NEWuser10101001010PASS"

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

func TestUserPasswordEdit(test *testing.T) {
	userInfo := GetUserByName(userTestName)

	fmt.Printf("Editing user %d with a new pass!", userInfo.UserID)
	if hasChanged := EditUserPass(editUserTestPass, userInfo.UserID); !hasChanged {
		test.Fatalf("Cannot edit user %s password!", userInfo.UserName)
	} else {
		userInfo = GetUserByID(userInfo.UserID)
		if userNewPass := GetUserPass(userInfo.UserID); userNewPass == "NONE" || userNewPass == userTestPass {
			test.Fatalf("The pass didn't change at all, old pass: %s | new pass: %s", userTestPass, userNewPass)
		}
	}
}

func BenchmarkDBInsert(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		InsertUser(fmt.Sprintf("USER_%d_BENCH", i), "userpass", 0)
	}
}
