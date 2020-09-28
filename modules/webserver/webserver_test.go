package webserver

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

const serverIP string = "127.0.0.1"
const serverPort int = 8080

func TestPing(test *testing.T) {
	ts := httptest.NewUnstartedServerhttp(LoadWebServer())

	defer ts.Close()

	formatURL(ts, "/api/ping")
	res, resErr := http.Get(ts.URL)

	if resErr != nil {
		test.Fatal(resErr)
	}

	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		test.Fatal(err)
	}

	fmt.Printf("%s", greeting)
}
func TestLogin(test *testing.T) {

}

func formatURL(ts *httptest.Server, route string) {
	ts.URL = fmt.Sprintf("http://%s:%d%s", serverIP, serverPort, route)
}
