package middleware

import (
	"net/http"
	"strings"
)

type CorsMiddleware struct {
	AllowedOrigins []string
	AllowedHeaders string
}

// Provides a middleware that allows Cross-Origin-Requests.
func (cm *CorsMiddleware) Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if contains(cm.AllowedOrigins, r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		}
		r.Header.Del("Origin")
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, DELETE")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", cm.AllowedHeaders)

		next.ServeHTTP(w, r)
	})
}

func contains(s []string, search string) bool {
	contains := false

	for entry := range s {
		if strings.Contains(search, s[entry]) {
			contains = true
		}
	}

	return contains
}
