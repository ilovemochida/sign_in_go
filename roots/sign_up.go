package roots

import (
	"net/http"
	"text/template"
)

type Page struct {
	Title string
}

func SignUp(w http.ResponseWriter, r *http.Request) {
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
