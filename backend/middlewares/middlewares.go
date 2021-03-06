package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/jpoly1219/go-leetcode/backend/auth"

	"github.com/golang-jwt/jwt"
)

type key string

var authKey key = "authkey"

func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth.HandleCors(w, r)
		if r.Method == "OPTIONS" {
			return
		}
		fmt.Println(r.Header.Get("Authorization"))
		authHeader := strings.Split(r.Header.Get("Authorization"), " ")
		// fmt.Println("0: ", authHeader[0], "1:", authHeader[1])
		if len(authHeader) != 2 {
			fmt.Println("bad request: no token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("ACCESSSECRETKEY")), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), authKey, claims)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				fmt.Println("invalid claims:", err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		}
	})
}
