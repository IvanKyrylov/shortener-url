package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/IvanKyrylov/shortener-url/internal/cache"
)

var storage = cache.NewStorage()

func responseHandler(h func(http.ResponseWriter, *http.Request) (interface{}, int, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, status, err := h(w, r)
		if err != nil {
			data = err.Error()
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		err = json.NewEncoder(w).Encode(response{Data: data, Success: err == nil})
		if err != nil {
			log.Printf("could not encode response to output: %v", err)
		}
	}
}

func cached(duration string, handler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		content := storage.Get(r.RequestURI)
		if content != nil {
			log.Print("Cache Hit!\n")
			w.Write(content)
		} else {
			c := httptest.NewRecorder()
			handler(c, r)

			for k, v := range c.HeaderMap {
				w.Header()[k] = v
			}

			w.WriteHeader(c.Code)
			content := c.Body.Bytes()

			if d, err := time.ParseDuration(duration); err == nil {
				log.Printf("New page cached: %s for %s\n", r.RequestURI, duration)
				storage.Set(r.RequestURI, content, d)
			} else {
				log.Printf("Page not cached. err: %s\n", err)
			}

			w.Write(content)
		}

	}
}
