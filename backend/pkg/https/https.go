package https

import (
	"backend/pkg/logger"
	"encoding/json"
	"io"
	"net/http"
)

type contextKey struct {
	name string
}

type ErrorBody struct {
	Code    int
	Message string
}

type JsonBody struct {
	Code int
	Data interface{}
}

var StatusCtxKey = &contextKey{"Status"}

func ResponseError(w http.ResponseWriter, r *http.Request, statusCode int, v string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	byteBody, err := json.Marshal(ErrorBody{
		Code:    -1,
		Message: v,
	})
	if err != nil {
		logger.Trace(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(byteBody)
}

func ResponseJSON(w http.ResponseWriter, r *http.Request, statusCode int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	byteBody, err := json.Marshal(JsonBody{
		Code: 0,
		Data: v,
	})
	if err != nil {
		logger.Trace(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(byteBody)
}

func ResponseMsg(w http.ResponseWriter, r *http.Request, statusCode int, v string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	byteBody, err := json.Marshal(ErrorBody{
		Code:    0,
		Message: v,
	})
	if err != nil {
		logger.Trace(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(byteBody)
}

func GetBody[T any](r *http.Request) (T, error) {
	var data T
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func GetQuery[T any](r *http.Request) (T, error) {
	var data T
	queryParams := r.URL.Query()
	
	err := json.Unmarshal([]byte(queryParams.Encode()), data)
	if err != nil {
		return data, err
	}
	return data, nil
}
