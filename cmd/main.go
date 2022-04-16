package main

import (
	"net/http"
	"simple-todo-api/pkg"
)

// server を立ち上げてリクエスト処理へのルーティングを行う
func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/todos", pkg.HandleTodosRequest)
	server.ListenAndServe()
}
