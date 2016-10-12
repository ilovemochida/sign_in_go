package main

import (
	"github.com/gorilla/mux"
	"maegari/roots"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", roots.ViewHandler)
	r.HandleFunc("/sign_up", roots.SignUpHandler)
	r.HandleFunc("/login", roots.Login)
	r.NotFoundHandler = http.HandlerFunc(roots.NotFoundHandler)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
