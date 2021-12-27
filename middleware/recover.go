package middleware

import (
	"log"
	"net/http"
)

func (m Middleware) RecoverHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("recover from panic %v", err)
				http.Error(w, http.StatusText(500), 500)

			}
		}()
	}
	return http.HandlerFunc(fn)
}

