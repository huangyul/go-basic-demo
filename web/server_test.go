package web

import (
	"testing"
)

func TestServer(t *testing.T) {
	var h Server = &HTTPServer{}

	// 用法一：完全给http包
	//http.ListenAndServe(":8081", h)
	//http.ListenAndServeTLS(":8082", "", "", h)

	// 自己手动启动
	h.Start(":8083")
}
