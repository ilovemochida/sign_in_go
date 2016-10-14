package roots

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"text/template"
)

type signUpPage struct {
	Title string
	Miss  bool
}

func SignUp(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "root:rootpass@/gosample")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	r.ParseForm()
	if r.Method == "GET" {

		page := signUpPage{"sign up page", false}
		tmpl, err := template.ParseFiles("./view/signUp.html")
		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(w, page)
		if err != nil {
			panic(err)
		}

	} else {
		//ログインデータがリクエストされ、ログインのロジック判断が実行されます。
		un := r.FormValue("username")
		pw := r.FormValue("password")
		pw = Encrypt(pw)
		success := true

		fmt.Println("暗号化したパスワード: ", pw)

		rows, err := db.Query("SELECT * FROM users where name = ?", un)
		if err != nil {
			fmt.Println(err)
		}

		for rows.Next() {
			var name string
			var password string
			if err := rows.Scan(&name, &password); err != nil {
				fmt.Println(err)
			}
			if name != "" {
				if name == un {
					success = false
				}
			}
		}
		defer rows.Close()

		if success == true {
			fmt.Println("登録成功")

			query := "INSERT INTO users values(?, ?)"
			result, err := db.Exec(query, un, pw)
			if err != nil {
				log.Fatal("insert error: ", err)
			}
			fmt.Println(result)
			http.Redirect(w, r, "/login", http.StatusFound)
		} else {
			fmt.Println("登録失敗")
			page := signUpPage{"sign up page", true}
			tmpl, err := template.ParseFiles("./view/signUp.html")
			if err != nil {
				panic(err)
			}

			err = tmpl.Execute(w, page)
			if err != nil {
				panic(err)
			}
		}
	}
}
