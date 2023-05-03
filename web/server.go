package web

import (
	"net"
	"net/http"
)

// 确保一定实现结构体实现了接口
var _ Server = &HTTPServer{}

type HandleFunc func()

type Server interface {
	http.Handler
	Start(addr string) error

	// 增加路由注册
	addRoute(method string, path string, handleFunc HandleFunc)
}

type HTTPServer struct {
}

func (h *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		Req:  request,
		Resp: writer,
	}
	h.serve(ctx)
}

func (h *HTTPServer) serve(ctx *Context) {
}

func (h *HTTPServer) Start(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return http.Serve(l, h)
}

func (h *HTTPServer) addRoute(method string, path string, handleFunc HandleFunc) {
}

func (h *HTTPServer) Get(path string, handleFunc HandleFunc) {
	h.addRoute(http.MethodGet, path, handleFunc)
}

func (h *HTTPServer) Post(path string, handleFunc HandleFunc) {
	h.addRoute(http.MethodPost, path, handleFunc)

}
