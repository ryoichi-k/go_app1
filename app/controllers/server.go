package controllers

import (
	"fmt"
	"golang_app/app_1/config"
	"html/template"
	"net/http"
)

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

// func session(writer http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
// 	//サーバーのリクエストからクッキーを取得
// 	cookie, err := request.Cookie("_cookie")
// 	if err == nil {
// 		sess = models.Session{UUID: cookie.Value} //sessionがあるかを探す
// 		// if ok, _ := sess.CheckSession(); !ok {
// 		// 	// err = errors.New("Invalid session")
// 		// }
// 	}
// 	return
// }

//webサーバに接続
func StartMainServer() error {
	//css,js読み込み処理
	fmt.Println("valid1")
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/todos", index)

	//localhost:8080接続
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
