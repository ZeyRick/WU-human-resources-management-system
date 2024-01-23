package helper

import (
	"backend/pkg/https"
	"backend/pkg/logger"
	"net/http"
)

func UnexpectedError(w http.ResponseWriter, r *http.Request, err error) {
	logger.Trace(err)
	https.ResponseError(w, r, http.StatusInternalServerError, "Something went wrong")
}
