package web

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	var h = &HTTPServer{}
	h.AddRoute(http.MethodGet, "/user", func(ctx Context) {})
	// 用法一：完全给http包q
	//http.ListenAndServe(":8081", h)
	//http.ListenAndServeTLS(":8082", "", "", h)

	// 自己手动启动
	h.Start(":8083")
}
