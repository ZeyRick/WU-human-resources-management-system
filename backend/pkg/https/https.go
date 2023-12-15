package https

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)
type contextKey struct {
	name string
}

var StatusCtxKey = &contextKey{"Status"}

func ResponseJSON(w http.ResponseWriter, r *http.Request, statusCode int,v interface{}) {
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
	w.WriteHeader(statusCode)
	w.Write(buf.Bytes()) //nolint:errcheck
}

func ResponseText(w http.ResponseWriter, r *http.Request, statusCode int,v string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if status, ok := r.Context().Value(StatusCtxKey).(int); ok {
		w.WriteHeader(status)
	}
	w.WriteHeader(statusCode)
	w.Write([]byte(v)) //nolint:errcheck
}

func GetBody [T any] (r *http.Request) (T, error) {
	var data T;
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(body,  &data )
	if err != nil {
		return data, err
	}
	return data, nil
}