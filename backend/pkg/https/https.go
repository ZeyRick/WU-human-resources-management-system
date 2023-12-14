package https

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)
type contextKey struct {
	name string
}

var StatusCtxKey = &contextKey{"Status"}

func ResponseJSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if status, ok := r.Context().Value(StatusCtxKey).(int); ok {
		w.WriteHeader(status)
	}
	w.Write(buf.Bytes()) //nolint:errcheck
}

func ResponseText(w http.ResponseWriter, r *http.Request, v string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if status, ok := r.Context().Value(StatusCtxKey).(int); ok {
		w.WriteHeader(status)
	}
	w.Write([]byte(v)) //nolint:errcheck
}

func GetBody [T any] (r *http.Request) (T, error) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var data T;
	err := decoder.Decode(&data);
	    if err != nil {
        var syntaxError *json.SyntaxError
        var unmarshalTypeError *json.UnmarshalTypeError
		var msg string
        switch {
        case errors.As(err, &syntaxError):
            msg = fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
        case errors.Is(err, io.ErrUnexpectedEOF):
            msg = fmt.Sprintf("Request body contains badly-formed JSON") 
        case errors.As(err, &unmarshalTypeError):
            msg = fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
        case strings.HasPrefix(err.Error(), "json: unknown field "):
            fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
            msg = fmt.Sprintf("Request body contains unknown field %s", fieldName)
        case errors.Is(err, io.EOF):
            msg = "Request body must not be empty"
        case err.Error() == "http: request body too large":
            msg = "Request body must not be larger than 1MB"
        default:
            msg = err.Error()
        }
        return data, errors.New(msg)
    } 
	err = decoder.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
        msg := "Request body must only contain a single JSON object"
        return data, errors.New(msg)
    }
	return data, nil
}