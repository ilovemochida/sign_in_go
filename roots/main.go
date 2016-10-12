package roots

import (
	"fmt"
	"net/http"
	"text/template"
)

type Page struct {
	Title string
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
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

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		//ログインデータがリクエストされ、ログインのロジック判断が実行されます。
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
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
