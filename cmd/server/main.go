package main

import (
	"database/sql"
	"log"
	"net/http"
	"simple-todo-api/pkg"

	// postgres ドライバ
	_ "github.com/lib/pq"
)

var Db *sql.DB

// init = パッケージの初期化処理などを行う main.goよりも先に実行される
func init() {
	var err error
	Db, err = sql.Open("postgres", "host=db user=simpleTodoDb password=password dbname=simpleTodoDb sslmode=disable")
	log.Println(&Db)
	if err != nil {
		log.Fatal(err)
	}
}

// メイン関数
func main() {

	server := http.Server{
		Addr: ":8080",
	}

	log.Print("start server")

	http.HandleFunc("/todos", pkg.HandleTodosRequest)
	server.ListenAndServe()
}
