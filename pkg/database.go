package pkg

import (
	"database/sql"
	"log"

	// postgres ドライバ
	_ "github.com/lib/pq"
)

var Db *sql.DB

// init = パッケージの初期化処理などを行う main.goよりも先に実行される
func NewDb() error {
	var err error
	Db, err = sql.Open("postgres", "host=db user=simpleTodoDb password=password dbname=simpleTodoDb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return err
}
