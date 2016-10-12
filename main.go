package main

import (
	"io"
	"net/http"
)

func infoHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `<h1>Hello, World!</h1>`)
}

func main() {
	http.HandleFunc("/", infoHandler)
	http.ListenAndServe(":8080", nil)
}
