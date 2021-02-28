package auth

import "net/http"

type ServiceAuth interface {
	GenerateToken() (string, error)
	ValidToken(r *http.Request) error
}
