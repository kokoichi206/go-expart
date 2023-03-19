package main

import "net/http"

func firstServer() {
	// static 配下のフォルダが、/public/ パスでサーブされる
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./static"))))
	// http://localhost:8004/public/logo.png
	http.ListenAndServe(":8004", nil)
}
