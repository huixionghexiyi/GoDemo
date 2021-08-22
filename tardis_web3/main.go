package main

import (
	"fmt"
	"log"
	"net/http"
	"tardis_web2/context"
	web "tardis_web2/server"
	"tardis_web2/sign"
)

func main() {
	server := web.NewSdkHttpServer("test_web")
	//server.Route("/", handler)
	server.Route("GET", "/user", user)
	server.Route("GET", "/sign", signUp)
	log.Fatal(server.Start("8080"))

}

func user(c context.AbstractContext) {
	c.WriteJson(http.StatusCreated, "Hi, this is home page")
}

func handler(c context.AbstractContext) {
	c.WriteJson(200, "Hi there, I love %s!")
}

func signUp(c context.AbstractContext) {
	req := &sign.SignUpReq{}
	err := c.ReadJson(req)

	if err != nil {
		resp := &context.CommonResponse{
			BizCode: 4,
			Msg:     fmt.Sprintf("invalid request: %v", err),
		}
		c.WriteJson(4, resp)
		return
	}

	_ = c.WriteJson(4, &context.CommonResponse{
		BizCode: 2,
		Data:    123,
	})
}
