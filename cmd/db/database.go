package db

import (
	"database/sql"
	"fmt"
	"log"
)

var Db *sql.DB

// 一旦 DB を作成し、そこから情報を取得するようにする
//　ゆくゆくは ID をもとに DB を取得 / 作成するようにしたい

// init = パッケージの初期化処理などを行う main.goよりも先に実行される
func init() {
	var err error
	//DSNの設定
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"simple-todo-app", "simple-todo-password", "simple-todo-api:3306", "todo")
	Db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer Db.Close()
}
