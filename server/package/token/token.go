package token

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

const (
	SigningKey    = "secret"
	SigningMethod = "HS256"
)

type Claims struct {
	Userid   uint   `json:"userid"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func New(userid uint, username string, expire int64) string {
	claims := Claims{
		userid,
		username,
		jwt.StandardClaims{
			ExpiresAt: expire,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte("secret"))

	return signedToken
}

func Parse(r *http.Request) (jwt.MapClaims, error) {
	authHeader := r.Header.Get("Authorization")
	authHeaderParts := strings.Split(authHeader, " ")

	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return nil, errors.New("Authorization is required.")
	}

	token, _ := jwt.Parse(authHeaderParts[1], func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if !token.Valid {
		return nil, errors.New("Token is invalid.")
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	return claims, nil
}
