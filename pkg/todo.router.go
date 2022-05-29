package pkg

import (
	"database/sql"
	"fmt"
	"log"
)

type Todos struct {
	id         int
	name       string
	todo       string
	created_at string
}

// GetTodos DB からデータを全件取得して一覧を返す
func GetTodos(db *sql.DB) {
	fmt.Println("GET")
	rows, err := db.Query("SELECT * FROM simpleTodoDb")
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var t Todos
		if err := rows.Scan(&t.id, &t.name, &t.todo, &t.created_at); err != nil {
			log.Println(err)
		}
		fmt.Printf("[%d]  Name:%s TODO:%s 作成日時:%s\n", t.id, t.name, t.todo, t.created_at)
	}
}

// PostTodo クライアントから送られてきたデータをもとに DB に追加
func PostTodo() {
	fmt.Println("POST")
}

// PutTodo クライアントから送られてきたデータをもとに DB を更新する
func PutTodo() {
	fmt.Println("PUT")
}

// DeleteTodo 指定データを DB から削除する
func DeleteTodo() {
	fmt.Println("DELETE")
}
