package main

import (
	"net/http"
	"fmt"
	"golang.org/x/net/html"
)

const (
	port = ":3000"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func acceptData(w http.ResponseWriter, r *http.Request) { 
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/acceptData", acceptData)
	http.ListenAndServe(port, nil)
}
