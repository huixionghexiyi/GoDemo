package main

import (
	"net/http"
	"strings"
)

type TreeHandler struct {
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

// findMatchChild2 增加通配符匹配
func (n *node) findMatchChild2(path string) (*node, bool) {
	panic("TODO me")
}

// findMatchChild2 增加变量匹配
func (n *node) findMatchChild3(path string) (*node, bool) {
	panic("TODO me")
}

// ServeHTTP 可以抽出一个 findRouter 方法
func (h *TreeHandler) ServeHTTP(c *Context) {

	router, found := h.findRouter(c.R.URL.Path)
	if !found {
		c.W.WriteHeader(http.StatusNotFound)
		c.W.Write([]byte("Not Found"))
		return
	}
	router.handler(c)
}

func (h *TreeHandler) findRouter(path string) (*node, bool) {
	paths := strings.Split(strings.Trim(path, "/"), "/")
	curr := h.root
	for _, p := range paths {
		matchChild, found := curr.findMatchChild(p)
		if !found {
			return nil, false
		}
		curr = matchChild
	}
	return curr, true
}

// Route 路由
func (h *TreeHandler) Route(method, pattern string, handlerFunc handlerFunc) {
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

// createSubTree 创建子树
func (h *TreeHandler) createSubTree(root *node, paths []string, f handlerFunc) {
	curr := root

	// 遍历传进来的参数，不停的创建子节点，直到创建完成后，将 f 赋值给 curr.handler
	for _, path := range paths {
		nn := newNode(path)
		curr.children = append(curr.children, nn)
		curr = nn
	}
	curr.handler = f
}

// newNode 新建一个节点
func newNode(path string) *node {
	return &node{
		path:     path,
		children: make([]*node, 0, 2),
	}
}

func NewTreeHandler() *TreeHandler {
	return &TreeHandler{
		root: &node{},
	}
}
