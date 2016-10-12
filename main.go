package main

import (
	"net/http"
	"text/template"
)

type Page struct {
	Title string
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{"Hello World."}
	tmpl, err := template.ParseFiles("./view/index.html") // ParseFilesを使う
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
