package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandlerFunc)
	http.ListenAndServe(":8081", nil)
}

func indexHandlerFunc(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "<p>go-demo1</p>")
}
