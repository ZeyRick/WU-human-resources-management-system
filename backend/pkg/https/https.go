package https

import (
	"backend/adapters/dtos"
	"backend/pkg/encrypt"
	"backend/pkg/logger"
	"backend/pkg/variable"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
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
	Code int         `json:"code"`
	Res  interface{} `json:"res"`
}

var StatusCtxKey = &contextKey{"Status"}
var decoder = schema.NewDecoder()

func ResponseError(w http.ResponseWriter, r *http.Request, statusCode int, v string) {
	byteBody := ErrorBody{
		Code: -1,
		Msg:  v,
	}
	w.WriteHeader(statusCode)
	render.JSON(w, r, byteBody)
}

func ResponseJSON(w http.ResponseWriter, r *http.Request, statusCode int, v interface{}) {
	jsonStr, err := json.Marshal(v)
	if err != nil {
		logger.Trace(err)
	}
	encrypted, err := encrypt.AesCbcEncrypt(string(jsonStr), []byte(os.Getenv("AES_KEY")))
	if err != nil {
		logger.Trace(err)
	}
	byteBody := JsonBody{
		Code: 0,
		Res:  encrypted,
	}
	w.WriteHeader(statusCode)
	render.JSON(w, r, byteBody)
}

func ResponseMsg(w http.ResponseWriter, r *http.Request, statusCode int, v string) {
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

func GetParamsStr(r *http.Request, key string) (string, error) {
	str := chi.URLParam(r, key)
	return str, nil
}

func GetParamsID(r *http.Request, key string) (*int, error) {
	idStr := chi.URLParam(r, key)
	if idStr == "" {
		return nil, nil
	}
	userId, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}
	return &userId, nil
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

func GetPaginationWithType[T any](r *http.Request) (dtos.PageOpt, T, error) {
	decoder.IgnoreUnknownKeys(true)
	var pageOpt dtos.PageOpt
	var filter T
	err := decoder.Decode(&pageOpt, r.URL.Query())
	if err != nil {
		return pageOpt, filter, err
	}
	err = decoder.Decode(&filter, r.URL.Query())
	if err != nil {
		return pageOpt, filter, err
	}
	if pageOpt.Page == nil {
		pageOpt.Page = variable.Create[int64](1)
	}
	if pageOpt.Size == nil {
		pageOpt.Size = variable.Create[int64](10)
	}
	return pageOpt, filter, nil
}
func GetParamsStrV2(r *http.Request, key string) (string, error) {
	str := r.URL.Query().Get(key)
	if str == "" {
		return "", fmt.Errorf("query parameter '%s' not found", key)
	}
	return str, nil
}
