package web

// 用来支持对路由树的操作
// 代表路由树
type router struct {
	// http method 到路由树根节点
	trees map[string]*node
}

func newRouter() *router {
	return &router{
		trees: map[string]*node{},
	}
}

func (r *router) AddRoute(method string, path string, handleFunc HandleFunc) {

}

type node struct {
	// 路径
	path string
	// 子路径
	children map[string]*node
	// 处理业务的方法
	handler HandleFunc
}
