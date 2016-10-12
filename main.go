package main

import (
	"fmt"
	"github.com/gorilla/mux"
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

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", viewHandler)
	r.HandleFunc("/sign_up", signUpHandler)
	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
