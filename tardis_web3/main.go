package main

import (
	"fmt"
	"net/http"
	"tardis_web3/sign"
)

func main() {
	server := NewSdkHttpServer("test_web", MetricFilterBuilder)
	//server.Route("/", handler)
	server.Route(http.MethodGet, "/user", user)
	server.Route(http.MethodGet, "/sign", signUp)
	// 启动快速失败
	if err := server.Start("8080"); err != nil {
		panic(err)
	}

}

func user(c *Context) {
	c.WriteJson(http.StatusCreated, "Hi, this is home page")
}

func handler(c *Context) {
	c.WriteJson(200, "Hi there, I love %s!")
}

func signUp(c *Context) {
	req := &sign.SignUpReq{}
	err := c.ReadJson(req)

	if err != nil {
		resp := &CommonResponse{
			BizCode: 4,
			Msg:     fmt.Sprintf("invalid request: %v", err),
		}
		c.WriteJson(4, resp)
		return
	}

	_ = c.WriteJson(4, &CommonResponse{
		BizCode: 2,
		Data:    123,
	})
}
