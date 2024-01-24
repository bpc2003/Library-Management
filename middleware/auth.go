package middleware

import (
	"net/http"
	"os"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		TOK, _ := os.LookupEnv("MASTER_TOKEN")

		if r.Header.Get("authorization") != TOK {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("You're not Permitted to perform this action"))
			return
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
