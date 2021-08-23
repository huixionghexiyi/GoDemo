package main

type Routable interface {
	Route(method, pattern string, handlerFunc func(c *Context))
}
