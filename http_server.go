package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

type HttpHandler struct{}

func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	handler := HttpHandler{}

	http.ListenAndServe(":8090", handler)
}
