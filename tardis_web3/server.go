package server

import (
	"net/http"
	"tardis_web3/context"
	"tardis_web3/handler"
	"tardis_web3/route"
)

type Server interface {

	// Routable 设定一个路由，命中该路由的会执行handlerFunc的代码
	route.Routable

	// Start 启动我们的服务器
	Start(address string) error
}

// sdkHttpServer 这个是基于 net/http 这个包实现的 http server
type sdkHttpServer struct {
	// Name server 的名字，给个标记，日志输出的时候用得上
	Name string

	// handler 依赖于接口
	handler handler.Handler
}

func (s *sdkHttpServer) Route(method, pattern string, handlerFunc func(c context.AbstractContext)) {
	// 直接调用 handler 的 Route 方法
	s.handler.Route(method, pattern, handlerFunc)
}

func (s *sdkHttpServer) Start(address string) error {

	return http.ListenAndServe(":"+address, s.handler)
}

// NewSdkHttpServer 创建一个SdkHttpServer 对象
func NewSdkHttpServer(name string) Server {
	handler := &handler.MapHandler{
		Handlers: make(map[string]func(c context.AbstractContext)),
	}
	return &sdkHttpServer{
		Name:    name,
		handler: handler,
	}
}

var _ Server = &sdkHttpServer{}
