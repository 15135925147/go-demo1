package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandlerFunc)
	http.ListenAndServe(":8081", nil)

}

func indexHandlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "<p>go-demo1</p>")
		fmt.Fprintf(w, "当前请求url为：%v", r.URL.Path)
	} else if r.URL.Path == "/home" {
		fmt.Fprintf(w, "<p>home 目录</p>")
		fmt.Fprintf(w, "当前请求url为：%v", r.URL.Path)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}
