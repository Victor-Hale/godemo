package main

import (
	"fmt"
	"net/http"
)

type Engine struct {

}
func (engine*Engine)ServeHTTP(w http.ResponseWriter,req * http.Request){
	switch req.URL.Path {
	case "/":
		_, err := fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
		if err != nil {
			return
		}
	case "/hello":
		for k, v := range req.Header {
			_, err := fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
			if err != nil {
				return
			}
		}
	default:
		_, err := fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
		if err != nil {
			return
		}
	}
	}
func main() {

	engine := new(Engine)

	http.ListenAndServe(":8080",engine)
}
