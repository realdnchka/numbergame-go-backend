package utils

import (
	"log"
	"net/http"
)

func Logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[endpoint]: %s; [device]: %s", r.URL.Path, r.UserAgent())
		f(w, r)
	}
}