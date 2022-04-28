package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	SECRET  = "super_secret_token"
	api_key = "1234"
)

func main() {
	http.HandleFunc("/jwt", GetJWT)
	http.Handle("/api", ValidateJWT(HomeHandle))

	http.ListenAndServe(":3000", nil)
}

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "super secret area")
}

func GetJWT(w http.ResponseWriter, r *http.Request) {
	api := r.Header["Api"]

	if api != nil {
		if strings.EqualFold(api_key, api[0]) {
			token, err := createJWT()
			if err != nil {
				return
			}

			fmt.Fprint(w, token)
		}

	}

	return
}

func createJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenStr, err := token.SignedString([]byte(SECRET))
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return tokenStr, nil
}

func ValidateJWT(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenReq := r.Header["Token"]
		if tokenReq != nil {
			token, err := jwt.Parse(tokenReq[0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("not authorized"))
				}

				return []byte(SECRET), nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("not authorized: " + err.Error()))
			}

			if token.Valid {
				next(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("not authorized"))
		}
	})
}
