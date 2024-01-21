package loggers

import (
	"bytes"
	"net/http"
)

// https://qiita.com/nakaryooo/items/8d12d284c9caceb9c947
type InterceptWriter struct {
	http.ResponseWriter
	status int
	body   *bytes.Buffer
}

func (w *InterceptWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *InterceptWriter) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func NewInterceptWriter(w http.ResponseWriter) *InterceptWriter {
	return &InterceptWriter{
		body:           bytes.NewBufferString(""),
		ResponseWriter: w,
	}
}
