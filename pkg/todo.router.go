package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Todos struct {
	todoId     string
	name       string
	userId     string
	created_at string
	updated_at string
}

type EncodeTodo struct {
	todoId     string    `json:"id"`
	name       string    `json:"name"`
	userId     string    `json:"todo"`
	created_at time.Time `json:"due_date"`
	updated_at time.Time `json:"due_date"`
}

// GetTodos DB からデータを全件取得して一覧を返す
func GetTodos(db *sql.DB, w http.ResponseWriter) {
	rows, err := db.Query("SELECT * FROM todos WHERE user_id = 'testUser'")
	if err != nil {
		log.Println(err)
	}
	var todoResponses Todos
	for rows.Next() {
		if err := rows.Scan(&todoResponses.todoId, &todoResponses.name, &todoResponses.userId, &todoResponses.created_at, &todoResponses.updated_at); err != nil {
			log.Println(err)
		}
	}
	rows.Close()
	log.Println(todoResponses)

	output, _ := json.MarshalIndent(todoResponses, "", "\t\t")
	log.Println(output)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(output)

}

// PostTodo クライアントから送られてきたデータをもとに DB に追加
func AddTodo(db *sql.DB) {
	fmt.Println("POST")
	fmt.Println(db)
}

// PutTodo クライアントから送られてきたデータをもとに DB を更新する
func EditTodo(db *sql.DB) {
	fmt.Println("PUT")
	fmt.Println(db)
}

// DeleteTodo 指定データを DB から削除する
func DeleteTodo(db *sql.DB) {
	fmt.Println("DELETE")
	fmt.Println(db)
}
