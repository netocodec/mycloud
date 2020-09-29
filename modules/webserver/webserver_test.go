package webserver

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"../db"
)

const userTest string = "userTest001"
const userPassTest string = "userPass001"

func TestLogin(test *testing.T) {
	prepareDB(test)
	router := LoadWebServer()
	reqJson := []byte(fmt.Sprintf(`{ "user": "%s", "pass":"%s" }`, userTest, userPassTest))
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(reqJson))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	responseCode := w.Code
	log.Printf("Request with status code %d", responseCode)

	if responseCode != http.StatusOK {
		test.Errorf("Status not returning OK, status returned: %d", responseCode)
	}
}

func prepareDB(test *testing.T) {
	fmt.Println("Preparing DB for API Testing...")

	if isRegistered := db.InsertUser(userTest, userPassTest, 0); !isRegistered {
		test.Fatal("Cannot register user into the database!")
	}
}
