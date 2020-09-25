package db

import (
	"fmt"
	"log"
	"testing"
)

const userTestName string = "usertest001"

func TestDBInit(test *testing.T) {
	InsertUser(userTestName, "userpass", 0)

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

func BenchmarkDBInsert(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		InsertUser(fmt.Sprintf("USER_%d_BENCH", i), "userpass", 0)
	}
}
