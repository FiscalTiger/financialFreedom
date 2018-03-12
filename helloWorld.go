package main

import (
	"net/http"
	"fmt"
	"bytes"
	"io"
)

const (
	port = ":3000"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func acceptData(w http.ResponseWriter, response *http.Request) {
	newStr :=  readCloserToString(response.Body)
	fmt.Fprintf(w, newStr)
}

func readCloserToString(readCloser io.ReadCloser) (retVal string){
	buf := new(bytes.Buffer)
  buf.ReadFrom(readCloser)
	retVal = buf.String()
  return
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/acceptData", acceptData)
	http.ListenAndServe(port, nil)
}
