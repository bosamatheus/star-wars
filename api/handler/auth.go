package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bosamatheus/star-wars/api/presenter"
	"github.com/bosamatheus/star-wars/pkg/auth"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func MakeAuthHandlers(r *mux.Router, n negroni.Negroni, service auth.ServiceAuth) {
	r.Handle("/api/v1/auth/token", n.With(
		negroni.Wrap(generateToken(service)),
	)).Methods("GET", "OPTIONS").Name("generateToken")
}

func generateToken(service auth.ServiceAuth) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		clientID := r.Header.Get("client_id")
		if clientID == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad request"))
			return
		}
		token, err := service.GenerateToken()
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("Error: " + err.Error()))
			return
		}

		toJ := &presenter.Auth{
			Token:      token,
			Authorized: true,
		}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error: " + err.Error()))
		}
	})
}
