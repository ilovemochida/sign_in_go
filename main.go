package main

import (
	"net/http"
	"text/template"
)

type Page struct {
	Title string
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{"sign in page"}
	tmpl, err := template.ParseFiles("./view/index.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{"sign up page"}
	tmpl, err := template.ParseFiles("./view/signUp.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/sign_up", signUpHandler)
	http.ListenAndServe(":8080", nil)
}
