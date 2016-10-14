package roots

import (
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/scrypt"
	"net/http"
	"text/template"
)

var (
	commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
)

type loginPage struct {
	Title string
	Miss  bool
}

func Encrypt(text string) string {
	b := []byte("mochidaCrypt")

	//ソルト原文
	salt := base64.StdEncoding.EncodeToString(b)

	converted, _ := scrypt.Key([]byte(text), []byte(salt), 16384, 8, 1, 16)
	return hex.EncodeToString(converted[:])
}

func Login(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "root:rootpass@/gosample")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	r.ParseForm()
	if r.Method == "GET" {

		page := loginPage{"sign in page", false}
		tmpl, err := template.ParseFiles("./view/login.html")
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
		success := false

		fmt.Println("暗号化したパスワード: ", pw)

		rows, err := db.Query("SELECT * FROM users where name = ?", un)
		defer rows.Close()
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
				if name == un && pw == password {
					success = true
				}
			}
		}
		if success == true {
			fmt.Println("ログイン成功")
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			fmt.Println("ログイン失敗")
			page := loginPage{"sign in page", true}
			tmpl, err := template.ParseFiles("./view/login.html")
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
