package ping

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPing(test *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	Ping(c)

	if w.Code != http.StatusOK {
		test.Errorf("Status code not valid! (Status Code:: %d)", w.Code)
	}

	var got gin.H
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		test.Fatal(err)
	}

	if got["message"] != "pong" {
		test.Errorf("Message response error: %s", got["message"])
	}
}
