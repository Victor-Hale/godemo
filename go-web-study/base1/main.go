package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloworld(w http.ResponseWriter , req * http.Request)  {
	_, err := fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	if err != nil {
		return
	}
	for k , v :=range req.Header{
		_, err := fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		if err != nil {
			return
		}
	}
}

func main() {
   http.HandleFunc("/",helloworld)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		 log.Fatal("ListenAndServe: ", err)
	}
}
