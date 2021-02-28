package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/bosamatheus/star-wars/config"
	"github.com/dgrijalva/jwt-go"
)

type ServiceJWT struct {
	clientID string
}

func NewAuthService(clientID string) (*ServiceJWT, error) {
	s := &ServiceJWT{
		clientID: clientID,
	}
	return s, nil
}

func (s *ServiceJWT) GenerateToken() (string, error) {
	claims := jwt.MapClaims{}
	claims["client_id"] = s.clientID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	claims["authorized"] = true
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.API_SECRET))
}

func (s *ServiceJWT) ValidToken(r *http.Request) error {
	tokenString := extractToken(r)
	_, err := parseJWT(tokenString, config.API_SECRET)
	if err != nil {
		return err
	}
	return nil
}

func extractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func parseJWT(tokenString string, APISecret string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("An error has occurred")
		}
		return []byte(APISecret), nil
	})
}
