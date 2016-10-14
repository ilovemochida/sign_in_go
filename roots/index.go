package roots

import (
	"net/http"
	"text/template"
)

type Page struct {
	Title string
}

func Index(w http.ResponseWriter, r *http.Request) {

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
