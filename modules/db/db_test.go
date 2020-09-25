package db

import (
	"os"
	"testing"
)

func TestDBInit(test *testing.T) {
	InsertUser("usertest", "userpass", 0)

	if users := GetAllUsers(); len(users) == 0 {
		test.Error("There is no user to show!")
	}

	os.Remove(DbFilename)
}
