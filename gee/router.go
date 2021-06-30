package gee

import "log"

type router struct {
	handlers map[string]HandleFunc
}

func newRouter() *router  {
	return &router{handlers: make(map[string]HandleFunc)}
}

func (r *router) addRoute(method string, pattern string, handler HandleFunc)  {
	log.Printf("Route %4s - %s", method, pattern)   // todo %4s
}