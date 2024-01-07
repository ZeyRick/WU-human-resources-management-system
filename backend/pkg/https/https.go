package https

import (
	"backend/adapters/dtos"
	"backend/pkg/logger"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/schema"
)

type contextKey struct {
	name string
}

type ErrorBody struct {
	Code int
	Msg  string
}

type JsonBody struct {
	Code int
	Data  interface{}
}

var StatusCtxKey = &contextKey{"Status"}
var decoder = schema.NewDecoder()


func ResponseError(w http.ResponseWriter, r *http.Request, statusCode int, v string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	byteBody, err := json.Marshal(ErrorBody{
		Code: -1,
		Msg:  v,
	})
	if err != nil {
		logger.Trace(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(`"Res":`))	
	w.Write(byteBody)
}

func ResponseJSON(w http.ResponseWriter, r *http.Request, statusCode int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	byteBody, err := json.Marshal(JsonBody{
		Code: 0,
		Data:  v,
	})
	if err != nil {
		logger.Trace(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(`"Res":`))	
	_, err = w.Write(byteBody)
	if err != nil {
		logger.Trace(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ResponseMsg(w http.ResponseWriter, r *http.Request, statusCode int, v string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	byteBody, err := json.Marshal(ErrorBody{
		Code: 0,
		Msg:  v,
	})
	if err != nil {
		logger.Trace(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err != nil {
		logger.Trace(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(`"Res":`))	
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
	decoder.IgnoreUnknownKeys(true)
	var data T
	err := decoder.Decode(&data, r.URL.Query())
	if err != nil {
		return data, err
	}
	return data, nil
}

func GetPagination(r *http.Request) (dtos.PageOpt, error) {
	decoder.IgnoreUnknownKeys(true)
	var data dtos.PageOpt
	err := decoder.Decode(&data, r.URL.Query())
	if err != nil {
		return data, err
	}
	return data, nil
}
