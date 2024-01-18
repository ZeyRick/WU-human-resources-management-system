package https

import (
	"backend/adapters/dtos"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/render"
	"github.com/gorilla/schema"
)

type contextKey struct {
	name string
}

type ErrorBody struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type JsonBody struct {
	Code int
	Data interface{}
}

var StatusCtxKey = &contextKey{"Status"}
var decoder = schema.NewDecoder()

func ResponseError(w http.ResponseWriter, r *http.Request, statusCode int, v string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	byteBody := ErrorBody{
		Code: -1,
		Msg:  v,
	}
	w.WriteHeader(statusCode)
	render.JSON(w, r, byteBody)
}

func ResponseJSON(w http.ResponseWriter, r *http.Request, statusCode int, v interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	byteBody := JsonBody{
		Code: 0,
		Data: v,
	}
	w.WriteHeader(statusCode)
	render.JSON(w, r, byteBody)
}

func ResponseMsg(w http.ResponseWriter, r *http.Request, statusCode int, v string) {
	// w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Del("Transfer-Encoding")
	byteBody := ErrorBody{
		Code: 0,
		Msg:  v,
	}

	w.WriteHeader(statusCode)
	render.JSON(w, r, byteBody)

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
