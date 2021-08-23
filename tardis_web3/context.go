package context

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type AbstractContext interface {
	ReadJson(data interface{}) error
	WriteJson(status int, data interface{}) error
}

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func (c *Context) ReadJson(req interface{}) error {
	body, err := ioutil.ReadAll(c.R.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, req)
}

func (c *Context) WriteJson(status int, data interface{}) error {
	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = c.W.Write(bs)
	if err != nil {
		return err
	}
	c.W.WriteHeader(status)
	return nil
}

func NewContext(w http.ResponseWriter, r *http.Request) AbstractContext {
	return &Context{
		W: w,
		R: r,
	}
}

var _ AbstractContext = &Context{}
