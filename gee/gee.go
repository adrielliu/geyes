package gee

import (
	"net/http"
)

type HandleFunc func(*Context)

type Engine struct {
	router *router
}

func (e *Engine)ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	c := newContext(w, req)
	e.router.handle(c)
}


func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) AddRoute(method string, pattern string, function HandleFunc)  {
	e.router.addRoute(method, pattern, function)
}

func (e * Engine ) Get(pattern string, function HandleFunc)  {
	e.AddRoute("GET", pattern, function)
}

func (e * Engine ) Post(pattern string, function HandleFunc)  {
	e.AddRoute("POST", pattern, function)
}

func (e * Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}