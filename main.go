package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("App is started;")
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(":8090", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello %s!", req.RemoteAddr)
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v, %v\n", name, h)
		}
	}
}
