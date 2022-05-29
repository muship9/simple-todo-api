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
	log.Println("Start DB")
	var err error
	Db, err = sql.Open("postgres", "user=admin password=admin dbname=simpleTodoDb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(Db)
}

// メイン関数
func main() {

	server := http.Server{
		Addr: ":8080",
	}

	log.Println("Start Server")

	http.HandleFunc("/todos", pkg.HandleTodosRequest)
	server.ListenAndServe()
}
