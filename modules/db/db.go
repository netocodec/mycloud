package db

import (
	"database/sql"
	"log"
	"os"

	dbInit "./dbinitquery"

	_ "github.com/mattn/go-sqlite3"
)

type UsersList struct {
	UserName,
	UserPassword string
	UserID,
	IsAdmin int
}

const dbFilename string = "mycloud.db"

var db *sql.DB

func init() {
	checkFile()

	log.Println("Init DB...")
	db, _ = sql.Open("sqlite3", dbFilename)
	initDatabase()
}

func InsertUser(username, password string, isAdmin int) bool {
	var result = true

	if statement, statErr := db.Prepare(dbInit.InsertUserQuery); statErr == nil {
		if _, execErr := statement.Exec(username, password, isAdmin); execErr != nil {
			result = false
			log.Fatalf("Error creating user: %s", execErr.Error())
		} else {
			log.Printf("User %s created with success! (Is Admin: %d)", username, isAdmin)
		}
	}

	return result
}

func GetAllUsers() []UsersList {
	var userResult []UsersList

	if queryResult, queryErr := db.Query(dbInit.SelectUsersQuery); queryErr == nil {
		defer queryResult.Close()

		for queryResult.Next() {
			var id int
			var isAdmin int
			var userName string

			queryResult.Scan(&id, &userName, &isAdmin)

			userResult = append(userResult, UsersList{
				UserID:       id,
				UserName:     userName,
				UserPassword: "NONE",
				IsAdmin:      isAdmin,
			})
		}
	}

	return userResult
}

func DeleteUserByID(id int) bool {
	var result = true

	if statement, statErr := db.Prepare(dbInit.DeleteUserQuery); statErr == nil {
		if _, execErr := statement.Exec(id); execErr != nil {
			result = false
			log.Fatalf("Error deleting user: %s", execErr.Error())
		} else {
			log.Printf("User %d deleted with success!", id)
		}
	}

	return result
}

func GetUserByName(username string) UsersList {
	var userResult UsersList

	if userQuery, userQueryErr := db.Query(dbInit.GetUserQuery, username); userQueryErr == nil {
		defer userQuery.Close()

		if userQuery.Next() {
			var id int
			var isAdmin int
			var userName string

			userQuery.Scan(&id, &userName, &isAdmin)

			userResult = UsersList{
				UserID:       id,
				UserName:     userName,
				UserPassword: "NONE",
				IsAdmin:      isAdmin,
			}
		}
	}

	return userResult
}

func CreateSharedLink(linkName, linkDirectory string, userID int) bool {
	result := true

	if sharedLink, sharedLinkErr := db.Prepare(dbInit.InsertSharedLinkQuery); sharedLinkErr == nil {
		if _, execErr := sharedLink.Exec(userID, linkName, linkDirectory); execErr != nil {
			result = false
		}
	}

	return result
}

func initDatabase() {
	if statement, statErr := db.Prepare(dbInit.InitUsersQuery); statErr == nil {
		statement.Exec()
		log.Println("Users table created with success!")
	} else {
		log.Printf("Cannot create Users table, error: %s", statErr.Error())
	}
}

func checkFile() {
	fileObj, err := os.OpenFile(dbFilename, os.O_WRONLY, 600)

	defer fileObj.Close()

	if err != nil {
		newFileObj, _ := os.Create(dbFilename)

		newFileObj.Close()
	}
}
