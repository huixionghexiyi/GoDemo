package main

import (
	"fmt"
	"log"
	"net/http"
	"tardis_web/context"
	web "tardis_web/server"
	"tardis_web/sign"
)

func main() {
	server := web.NewSdkHttpServer("test_web")
	server.Route("/", handler)
	server.Route("/user", user)
	server.Route("/sign", signUp)
	log.Fatal(server.Start("8080"))

}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, this is home page")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func signUp(w http.ResponseWriter, r *http.Request) {
	req := &sign.SignUpReq{}
	c := context.NewContext(w, r)
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

