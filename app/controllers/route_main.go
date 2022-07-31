package controllers

import (
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	// _, err := session(w, r)
	// if err != nil {
	//セッションが存在しないとき
	generateHTML(w, nil, "layout", "public_navbar", "top")
	// } else {
	// 	http.Redirect(w, r, "/todos", http.StatusFound)
	// }
}

//index.htmlに遷移するハンドラ
func index(w http.ResponseWriter, r *http.Request) {
	// _, err := session(w, r)
	// if err != nil {
	// 	http.Redirect(w, r, "/", http.StatusFound)
	// } else {
	generateHTML(w, nil, "layout", "private_navbar", "index")
	// }
}
