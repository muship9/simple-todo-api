package pkg

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Todos struct {
	id         int
	name       string
	todo       string
	created_at string
}

type EncodeTodo struct {
	id         int       `json:"id"`
	name       string    `json:"name"`
	todo       string    `json:"todo"`
	created_at time.Time `json:"due_date"`
}

// GetTodos DB からデータを全件取得して一覧を返す
func GetTodos(db *sql.DB, w http.ResponseWriter) {
	log.Println("GET")
	rows, err := db.Query("SELECT * FROM simpleTodoDb")
	if err != nil {
		log.Println(err)
	}
	var t Todos
	for rows.Next() {
		if err := rows.Scan(&t.id, &t.name, &t.todo, &t.created_at); err != nil {
			log.Println(err)
		}
		fmt.Printf("[%d]  Name:%s TODO:%s 作成日時:%s\n", t.id, t.name, t.todo, t.created_at)
	}
	rows.Close()

	// 取得されたデータを JSON に変換し、レスポンスとして返却する
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&t); err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, buf.String())

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
