package controllers

import (
	"golang_app/app_1/app/models"
	"log"
	"net/http"
)

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// _, err := session(w, r)
		// if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "signup")
		// } else {
		// 	http.Redirect(w, r, "/todos", http.StatusFound)
		// }
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Panicln(err)
		}
		//userのstructの各フィールドの値として受け取っている
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}
		//topページにリダイレクト
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	// _, err := session(w, r)
	// if err != nil {
	generateHTML(w, nil, "layout", "public_navbar", "login")
	// } else {
	// 	http.Redirect(w, r, "/todos", http.StatusFound)
	// }
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		//そのメアドのユーザーいない場合
		log.Println(err)
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	//パスワード照合
	if user.PassWord == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}
		//クッキーにセッションIDを入れる
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		//上記で作成したクッキーをセットする。
		http.SetCookie(w, &cookie)
		//照合成功→topページへ
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		//passが一致しない場合
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}
	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value}
		session.DeleteSessionByUUID()
	}
	http.Redirect(writer, request, "/login", http.StatusFound)
}
