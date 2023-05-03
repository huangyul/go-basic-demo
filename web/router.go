package web

type router struct {
	// 每一种method都有一棵树
	// http method => 路由树根节点
	trees map[string]*node
}

func (r *router) AddRoute(method string, path string, handleFunc HandleFunc) {

}

func newRouter() *router {
	return &router{
		trees: map[string]*node{},
	}
}

type node struct {
	path string
	// 子path到子节点
	children map[string]*node
	HandleFunc
}
