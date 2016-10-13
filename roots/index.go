package roots

import (
	"net/http"
	"text/template"
)

type Page struct {
	Title string
}

func Index(w http.ResponseWriter, r *http.Request) {

	value_post := r.FormValue("value_post")
	value_get := r.FormValue("value_get")

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
