package middleware

import (
	"net/http"

	"github.com/bosamatheus/star-wars/pkg/auth"
	"github.com/codegangsta/negroni"
)

func Auth(service auth.ServiceAuth) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		err := service.ValidToken(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}
