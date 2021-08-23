package main

import (
	"net/http"
)

// Handler 使用组合 扩展 http.Handler 方法
type Handler interface {
	ServeHTTP(c *Context)
	Routable
}

type MapHandler struct {
	Handlers map[string]func(c *Context)
}

// ServeHTTP 实现 Handler 接口
func (h *MapHandler) ServeHTTP(c *Context) {
	key := h.key(c.R.Method, c.R.URL.Path)

	// 如果找到路由就执行
	if handler, ok := h.Handlers[key]; ok {
		handler(c)
	} else {
		c.W.WriteHeader(http.StatusNotFound)
		c.W.Write([]byte("not any router match"))
	}
}

func (h *MapHandler) Route(method, pattern string, handlerFunc func(c *Context)) {
	key := h.key(method, pattern)
	h.Handlers[key] = handlerFunc
}

func (h *MapHandler) key(method, pattern string) string {
	return method + "&" + pattern
}

func NewMapHandler() Handler {
	return &MapHandler{
		Handlers: make(map[string]func(c *Context)),
	}
}

var _ Handler = &MapHandler{}
