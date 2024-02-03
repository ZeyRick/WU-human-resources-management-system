package middlewares

import (
	"backend/pkg/https"
	"backend/pkg/jwttoken"
	"backend/pkg/logger"
	"context"
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
		ok, userId := jwttoken.ValidateToken(w, r, reqToken)
		if ok {
			logger.Console("222")
			ctx := context.WithValue(r.Context(), "userId", userId)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		https.ResponseError(w, r, http.StatusUnauthorized, "Unauthorized")
	})
}
