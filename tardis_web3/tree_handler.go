package main

import (
	"net/http"
	"strings"
)

type HandlerBaseOnTree struct {
	root *node
}

type handlerFunc func(c *Context)

type node struct {
	path     string
	children []*node

	// 如果是叶子节点，那么匹配之后就可以调用该方法
	handler handlerFunc
}

// findMatchChild 查找 当前节点是否有指定节点
func (n *node) findMatchChild(path string) (*node, bool) {
	for _, child := range n.children {
		if child.path == path {
			return child, true
		}
	}
	return nil, false
}

// ServeHTTP 可以抽出一个 findRouter 方法
func (h *HandlerBaseOnTree) ServeHTTP(c *Context) {
	url := strings.Trim(c.R.URL.Path, "/")
	paths := strings.Split(url, "/")
	curr := h.root
	for _, path := range paths {
		child, found := curr.findMatchChild(path)

		// 在某一个父节点没找找到，直接 返回404
		if !found {
			c.W.WriteHeader(http.StatusNotFound)
			_, _ = c.W.Write([]byte("Not Found"))
			return
		}
		curr = child
	}

	// 如果路由的方法为 nil ，也相当于没有找到
	if curr.handler == nil {
		c.W.WriteHeader(http.StatusNotFound)
		c.W.Write([]byte("Not found"))
		return
	}

	curr.handler(c)

}

func (h *HandlerBaseOnTree) Route(method, pattern string, handlerFunc handlerFunc) {
	pattern = strings.Trim(pattern, "/")
	paths := strings.Split(pattern, "/")

	curr := h.root

	for index, path := range paths {
		matchChild, found := curr.findMatchChild(path)
		if found {
			curr = matchChild
		} else {
			h.createSubTree(curr, paths[index:], handlerFunc)
			break
		}
	}

	curr.handler = handlerFunc
}

func (h *HandlerBaseOnTree) createSubTree(root *node, paths []string, f handlerFunc) {
	curr := root

	// 遍历传进来的参数，不停的创建子节点，直到创建完成后，将 f 赋值给 curr.handler
	for _, path := range paths {
		nn := newNode(path)
		curr.children = append(curr.children, nn)
		curr = nn
	}
	curr.handler = f
}

func newNode(path string) *node {
	return &node{
		path:     path,
		children: make([]*node, 0, 2),
	}
}
