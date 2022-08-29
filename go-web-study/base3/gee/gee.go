package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request) //先做一个类型HandlerFunc 用于接受req和写入response

type Engine struct {
	//key 由请求方法和静态路由地址构成，例如GET-/，这样针对相同的路由，请求方法不同,可以映射不同的处理方法(Handler)，value 是用户映射的处理方法。
	router map[string]HandlerFunc
}

// New 创建一个new方法，new是engine结构体的构造函数，当使用 gee.new()方法时即可使用engine结构体的方法.实际上就是执行一个初始化map的操作
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET 定义了添加 GET 请求的方法
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 定义了添加 POST 请求的方法
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		_, err := fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
		if err != nil {
			return 
		}
	}
}
// Run 定义了启动服务的方法
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
