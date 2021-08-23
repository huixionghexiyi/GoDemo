package main

import (
	"net/http"
)

type Server interface {

	// Routable 设定一个路由，命中该路由的会执行handlerFunc的代码
	Routable

	// Start 启动我们的服务器
	Start(address string) error
}

// sdkHttpServer 这个是基于 net/http 这个包实现的 http server
type sdkHttpServer struct {
	// Name server 的名字，给个标记，日志输出的时候用得上
	Name string

	// handler 依赖于接口
	handler Handler

	//
	root Filter
}

func (s *sdkHttpServer) Route(method, pattern string, handlerFunc func(c *Context)) {
	// 直接调用 handler 的 Route 方法
	s.handler.Route(method, pattern, handlerFunc)
}

func (s *sdkHttpServer) Start(address string) error {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		context := NewContext(writer, request)
		s.root(context)
	})
	return http.ListenAndServe(":"+address, nil)
}

// NewSdkHttpServer 创建一个SdkHttpServer 对象
func NewSdkHttpServer(name string, builders ...FilterBuilder) Server {
	handler := NewMapHandler()

	// 因为是一个链，所以把最后的业务逻辑处理，也做为一环
	var root Filter = handler.ServeHTTP

	// 从后往前，把所有的filter串起来
	for i := len(builders) - 1; i >= 0; i-- {
		b := builders[i]
		root = b(root)
	}

	return &sdkHttpServer{
		Name:    name,
		handler: handler,
		root:    root,
	}

}

var _ Server = &sdkHttpServer{}
