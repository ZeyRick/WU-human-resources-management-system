package middlewares

import (
	"backend/pkg/encrypt"
	"backend/pkg/https"
	"backend/pkg/jwttoken"
	"backend/pkg/logger"
	"context"
	"encoding/base64"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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
			ctx := context.WithValue(r.Context(), "userId", userId)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		https.ResponseError(w, r, http.StatusUnauthorized, "Unauthorized")
	})
}

func DecryptMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				logger.Trace(err)
			}
			if string(body) != "" {
				encryptedData, _ := base64.StdEncoding.DecodeString(string(body))
				decrypted, err := encrypt.AesCbcDecrypt(encryptedData, []byte(os.Getenv("AES_KEY")))
				if err != nil {
					logger.Trace(err)
				}
				r.Body = ioutil.NopCloser(strings.NewReader(decrypted))
			}
		}
		next.ServeHTTP(w, r)
	})
}
