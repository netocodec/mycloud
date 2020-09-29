package webserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"../db"
)

const userTest string = "userTest001"
const userPassTest string = "userPass001"

type TokenLoginResponse struct {
	Token string `json:"token"`
}

type UserSettingsResponse struct {
	Username     string `json:"UserName"`
	UserPassword string `json:"UserPassword"`
	UserID       int    `json:"UserID"`
	IsAdmin      int    `json:"IsAdmin"`
}

func TestLogin(test *testing.T) {
	prepareDB(test)
	loginUser(test)
}

func TestGetSettings(test *testing.T) {
	userToken := loginUser(test)
	router := LoadWebServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/membership/settings", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", userToken))
	router.ServeHTTP(w, req)

	responseCode := w.Code
	if responseCode != http.StatusOK {
		test.Errorf("Status no returning OK, status returned: %d", responseCode)
	} else {
		var userSettings UserSettingsResponse

		json.Unmarshal(w.Body.Bytes(), &userSettings)

		if userSettings.Username != userTest {
			test.Errorf("User settings are wrong! (Username: %s)", userSettings.Username)
		}
	}
}

func TestChangePasswordSettings(test *testing.T) {
	userToken := loginUser(test)
	router := LoadWebServer()
	w := httptest.NewRecorder()
	bodyRequest := []byte(`{ "pass": "123456789", "passc": "123456789" }`)
	req, _ := http.NewRequest("POST", "/api/membership/settings", bytes.NewBuffer(bodyRequest))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", userToken))
	router.ServeHTTP(w, req)

	responseCode := w.Code
	if responseCode != http.StatusOK {
		test.Errorf("Status no returning OK, status returned: %d", responseCode)
	}
}

func prepareDB(test *testing.T) {
	fmt.Println("Preparing DB for API Testing...")

	if isRegistered := db.InsertUser(userTest, userPassTest, 0); !isRegistered {
		test.Fatal("Cannot register user into the database!")
	}
}

func loginUser(test *testing.T) string {
	var resultToken string = "NONE"
	router := LoadWebServer()
	reqJSON := []byte(fmt.Sprintf(`{ "user": "%s", "pass":"%s" }`, userTest, userPassTest))
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(reqJSON))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	responseCode := w.Code
	log.Printf("Request with status code %d", responseCode)

	if responseCode != http.StatusOK {
		test.Errorf("Status not returning OK, status returned: %d", responseCode)
	} else {
		var tokenResponse TokenLoginResponse
		if convertJSONErr := json.Unmarshal(w.Body.Bytes(), &tokenResponse); convertJSONErr == nil {
			resultToken = tokenResponse.Token
		}
	}

	return resultToken
}
