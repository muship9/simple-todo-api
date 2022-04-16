package main

import (
	"fmt"
	"net/http"
	"simple-todo-api/pkg"
)

// server を立ち上げてリクエスト処理へのルーティングを行う

func main() {
	fmt.Println("simple-todo-api")
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/todos", pkg.HandleTodosRequest)
	server.ListenAndServe()
}
