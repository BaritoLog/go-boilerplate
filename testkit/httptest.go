package testkit

import (
	"net/http"
	"net/http/httptest"
)

func NewHttpTestServer(statusCode int, body []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Write(body)
	}))

}
