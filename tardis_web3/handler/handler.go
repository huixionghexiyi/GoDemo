package handler

import (
	"net/http"
	"tardis_web2/context"
	"tardis_web2/route"
)

// Handler 使用组合 扩展 http.Handler 方法
type Handler interface {
	http.Handler
	route.Routable
}

type MapHandler struct {
	Handlers map[string]func(c context.AbstractContext)
}

// ServeHTTP 实现 Handler 接口
func (h *MapHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	key := h.Key(req.Method, req.URL.Path)

	if handler, ok := h.Handlers[key]; ok {
		c := context.NewContext(writer, req)
		handler(c)
	} else {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("not any router match"))
	}
}

func (h *MapHandler) Route(method, pattern string, handlerFunc func(c context.AbstractContext)) {
	key := h.Key(method, pattern)
	h.Handlers[key] = handlerFunc
}

func (h *MapHandler) Key(method, pattern string) string {
	return method + "&" + pattern
}

var _ Handler = &MapHandler{}
