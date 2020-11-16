package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/IvanKyrylov/shortener-url/config"
	"github.com/IvanKyrylov/shortener-url/store"
)

var h handler
var key string

func init() {
	configPath := "../../config/config.json"
	config, _ := config.FromConfigFile(configPath)

	svc, _ := store.New(config)
	h = handler{config.Options.Prefix, svc}

	key, _ = svc.Save("https://translate.google.com.ua/")

}

func TestShortener(t *testing.T) {
	req, err := http.NewRequest("POST", "/shortener", strings.NewReader(`{"url":"https://translate.google.com.ua/"}`))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(responseHandler(h.shortener))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestRedirect(t *testing.T) {
	req, err := http.NewRequest("GET", "/"+key, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.redirect)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusFound)
	}
}

func TestInfo(t *testing.T) {
	req, err := http.NewRequest("GET", "/info/"+key, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(responseHandler(h.info))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
