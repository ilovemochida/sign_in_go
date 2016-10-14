package main

import (
	"github.com/gorilla/mux"
	root "maegari/roots"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", root.Index)
	r.HandleFunc("/sign_up", root.SignUp)
	r.HandleFunc("/login", root.Login)
	r.NotFoundHandler = http.HandlerFunc(root.NotFound)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)

}
