package web

import (
	"net"
	"net/http"
)

// 小技巧，确保一定实现了Server接口
var _ Server = &HTTPServer{}

type HandleFunc func(ctx Context)

type Server interface {
	http.Handler
	Start(addr string) error

	// AddRoute 增加注册路由的方法
	// method 是HTTP方法
	// path 是路由
	// handleFunc 处理方法
	AddRoute(method string, path string, handleFunc HandleFunc)
}

type HTTPServer struct {
}

// ServeHTTP 处理请求的入口
func (h *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		Req: request,
		Res: writer,
	}
	// 接下来就是查找路由，并且执行业务逻辑
	h.server(ctx)
}

func (h *HTTPServer) server(ctx *Context) {

}

func (h *HTTPServer) Start(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// 这里可以加一些回调函数
	return http.Serve(l, h)
}

func (h *HTTPServer) AddRoute(method string, path string, handleFunc HandleFunc) {
	// 注册到路由树里面
}

func (h *HTTPServer) Get(path string, handleFunc HandleFunc) {
	h.AddRoute(http.MethodGet, path, handleFunc)
}

func (h *HTTPServer) Post(path string, handleFunc HandleFunc) {
	h.AddRoute(http.MethodPost, path, handleFunc)
}
