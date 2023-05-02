package web

import (
	"net"
	"net/http"
)

// 小技巧，确保一定实现了Server接口
var _ Server = &HTTPServer{}

type Server interface {
	http.Handler
	Start(addr string) error
}

type HTTPServer struct {
}

// ServeHTTP 处理请求的入口
func (h *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPServer) Start(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// 这里可以加一些回调函数
	return http.Serve(l, h)
}
