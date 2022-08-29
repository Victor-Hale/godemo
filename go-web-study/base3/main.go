package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {

	//type HandlerFunc func(http.ResponseWriter, *http.Request)
	//t := new(gee.Engine )
	//t.GET("/whys", func(w http.ResponseWriter, req *http.Request){
	//	_,err:=fmt.Fprintf(w,"我真牛逼！")
	//	if err != nil {
	//		return
	//	}
	//})
	//err := t.Run(":8080")
	//if err != nil {
	//	return
	//}

    r := gee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		_, err := fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
		if err != nil {
			return
		}
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			_, err := fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
			if err != nil {
				return
			}
		}
	})

	err := r.Run(":9999")
	if err != nil {
		return
	}
}
