package middlewares

import (
	"backend/pkg/https"
	"backend/pkg/jwttoken"
	"net/http"
	"strings"
)

func LoginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		if reqToken == "" {
			https.ResponseError(w, r, http.StatusUnauthorized, "Unauthorized")
			return
		}
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) != 2 {
			https.ResponseError(w, r, http.StatusUnauthorized, "Unauthorized")
			return
		}
		reqToken = splitToken[1]
		ok, _ := jwttoken.ValidateToken(w, r, reqToken)
		if ok {
			next.ServeHTTP(w, r)
			return
		}
		https.ResponseError(w, r, http.StatusUnauthorized, "Unauthorized")
	})
}
