package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{} //给map[string]interface{}起了一个别名gee.H
//Context 创建一个Context结构体
type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	// response info
	StatusCode int
}

//创建构造方法 调用newContext
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// PostForm 提供了访问PostForm参数的方法。
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query 提供了访问Query参数的方法。
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status 写入响应 code
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 设置resps中的 Header
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}
//String 提供了快速构造String响应的方法。
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	_, err := c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
	if err != nil {
		return 
	}
}

// JSON 提供了快速构造JSON响应的方法。
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data 提供了快速构造Data响应的方法。
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	_, err := c.Writer.Write(data)
	if err != nil {
		return 
	}
}

// HTML 提供了快速构造HTML响应的方法。
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	_, err := c.Writer.Write([]byte(html))
	if err != nil {
		return 
	}
}