package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/IvanKyrylov/shortener-url/store"
)

type handler struct {
	prefix string
	store  store.Service
}

type response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"response"`
}

func New(prefix string, store store.Service) http.Handler {
	mux := http.NewServeMux()
	h := handler{prefix, store}
	mux.HandleFunc("/shortener/", responseHandler(h.shortener))
	mux.HandleFunc("/", h.redirect)
	// mux.HandleFunc("/", cached("60s", h.redirect))
	return mux
}

func (h handler) shortener(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	if r.Method != http.MethodPost {
		return nil, http.StatusMethodNotAllowed, fmt.Errorf("method %s not allowed", r.Method)
	}

	var input struct{ URL string }
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("Unable to decode json request body: %v", err)
	}

	url := strings.TrimSpace(input.URL)
	if url == "" {
		return nil, http.StatusBadRequest, fmt.Errorf("Url is empty")
	}

	if !strings.Contains(url, "http") {
		url = "http://" + url
	}

	c, err := h.store.Save(url)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("Could not store in database: %v", err)
	}

	return h.prefix + c, http.StatusCreated, nil
}

func (h handler) redirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	code := r.URL.Path[len("/"):]

	url, err := h.store.Load(code)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("URL Not Found"))
		return
	}

	http.Redirect(w, r, string(url), http.StatusFound)
}
