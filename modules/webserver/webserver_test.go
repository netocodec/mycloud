package webserver

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(test *testing.T) {
	router := LoadWebServer()
	reqJson := []byte(`{ "user": "userTest001", "pass":"userPass001" }`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(reqJson))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
}
