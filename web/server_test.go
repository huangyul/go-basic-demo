package web

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	var h = &HTTPServer{}
	h.addRoute(http.MethodGet, "/user", func() {})
	h.Get("/getUser", func() {})

	h.Start(":8081")
}
