package gee

import (
	"fmt"
	"net/http"
)

type HandleFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	Router map[string]HandleFunc
}

func (e *Engine)ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := e.Router[key];ok{
		handler(w, req)
	}else{
		fmt.Fprintf(w, "404 NOT FOUND!!!", req.URL)
	}
}


func New() *Engine {
	return &Engine{Router: make(map[string]HandleFunc)}
}

func (e *Engine) AddRoute(method string, pattern string, function HandleFunc)  {
	//key := strings.Join([]string{method, pattern}, "-")
	//var buf bytes.Buffer
	//buf.WriteString(method)
	//buf.WriteString("-")
	//buf.WriteString(pattern)
	//key := buf.String()
	key := method + "-" + pattern
	e.Router[key] = function
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