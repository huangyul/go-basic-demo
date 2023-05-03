package web

import "strings"

type router struct {
	// 每一种method都有一棵树
	// http method => 路由树根节点
	trees map[string]*node
}

func (r *router) AddRoute(method string, path string, handleFunc HandleFunc) {
	// 首先找到树来
	root, ok := r.trees[method]

	if !ok {
		// 还没有根节点
		root = &node{
			path: "/",
		}
		r.trees[method] = root
	}
	path = path[1:]
	segs := strings.Split(path, "/")
	for _, seg := range segs {
		children := root.childOrCreate(seg)
		root = children
	}
	root.handler = handleFunc

}

func (n *node) childOrCreate(seg string) *node {
	if n.children == nil {
		n.children = map[string]*node{}
	}
	res, ok := n.children[seg]
	if !ok {
		res = &node{
			path: seg,
		}
		n.children[seg] = res

	}
	return res
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
	handler  HandleFunc
}
