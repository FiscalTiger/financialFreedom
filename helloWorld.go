package main

import (
	"net/http"
	"fmt"
	"bytes"
)

const (
	port = ":3000"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func acceptData(w http.ResponseWriter, response *http.Request) {
	buf := new(bytes.Buffer)
  buf.ReadFrom(response.Body)
  newStr := buf.String()

	fmt.Fprintf(w, newStr)
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/acceptData", acceptData)
	http.ListenAndServe(port, nil)
}
