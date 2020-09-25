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

const DbFilename string = "mycloud.db"

var db *sql.DB

func init() {
	checkFile()

	log.Println("Init DB...")
	db, _ = sql.Open("sqlite3", DbFilename)
	initDatabase()
}

func InsertUser(username, password string, isAdmin int) {
	if statement, statErr := db.Prepare(dbInit.InsertUserQuery); statErr == nil {
		if _, execErr := statement.Exec(username, password, isAdmin); execErr != nil {
			log.Fatalf("Error creating user: %s", execErr.Error())
		} else {
			log.Printf("User %s created with success! (Is Admin: %d)", username, isAdmin)
		}
	}
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

func initDatabase() {
	if statement, statErr := db.Prepare(dbInit.UsersQuery); statErr == nil {
		statement.Exec()
		log.Println("Users table created with success!")
	} else {
		log.Printf("Cannot create Users table, error: %s", statErr.Error())
	}
}

func checkFile() {
	fileObj, err := os.OpenFile(DbFilename, os.O_WRONLY, 600)

	defer fileObj.Close()

	if err != nil {
		newFileObj, _ := os.Create(DbFilename)

		newFileObj.Close()
	}
}
