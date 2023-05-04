package web

import (
	"fmt"
	"strings"
)

type router struct {
	// 每一种method都有一棵树
	// http method => 路由树根节点
	trees map[string]*node
}

func (r *router) AddRoute(method string, path string, handleFunc HandleFunc) {
	if path == "" {
		panic("web: the path cannot be empty")
	}
	if path[0] != '/' {
		panic("web: the path must start with '/'")
	}
	if path != "/" && path[len(path)-1] == '/' {
		panic("web: the path cannot end in ''/")
	}

	// 首先找到树来
	root, ok := r.trees[method]

	if !ok {
		// 还没有根节点
		root = &node{
			path: "/",
		}
		r.trees[method] = root
	}
	// root node special handling
	if path == "/" {
		if root.handler != nil {
			panic("duplicate route: /")
		}
		root.handler = handleFunc
		return
	}

	segs := strings.Split(path[1:], "/")
	for _, seg := range segs {
		if seg == "" {
			panic("web: path cannot have '//'")
		}
		children := root.childOrCreate(seg)
		root = children
	}
	if root.handler != nil {
		panic(fmt.Sprintf("duplicate route: %s", path))
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
