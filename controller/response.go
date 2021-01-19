package controller

import (
	"net/http"
)

// Response ...
type Response struct {
	http.ResponseWriter
	status int
	body   string
}

// NewResponse ...
func NewResponse(w http.ResponseWriter) *Response {
	return &Response{
		ResponseWriter: w,
	}
}

// Write ...
func (me *Response) Write(v []byte) (int, error) {
	me.body = string(v)
	return me.ResponseWriter.Write(v)
}

// WriteHeader ...
func (me *Response) WriteHeader(code int) {
	me.status = code
	me.ResponseWriter.WriteHeader(code)
}
