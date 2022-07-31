package main

import (
	"github.com/cagiti/go-tawerin/pkg/app"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var a = app.App{}

func TestApp(t *testing.T) {
	a = app.App{}
	a.Initialize()

	t.Run("homePageViaWeb", homePageViaWeb)
	t.Run("pingPageViaWeb", pingPageViaWeb)
}

func homePageViaWeb(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); !strings.Contains(body, "<title>Tawerin</title>") {
		t.Errorf("Expected a correct title. Got %s", body)
	}
}

func pingPageViaWeb(t *testing.T) {
	req, _ := http.NewRequest("GET", "/ping", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); !strings.Contains(body, "OK") {
		t.Errorf("Expected a correct body. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
