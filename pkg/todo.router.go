package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Todos struct {
	todoId     string
	title      string
	userId     string
	created_at time.Time
	updated_at time.Time
}

type EncodeTodo struct {
	TodoId string `json:"id"`
	Title  string `json:"title"`
	UserId string `json:"todo"`
}

type TodoRequest struct {
	Title  string `json:"title"`
	UserId string `json:"user_id"`
}

// GetTodos DB からデータを全件取得して一覧を返す
func GetTodos(db *sql.DB, w http.ResponseWriter) {

	// TODO : リファクタ -> DB からデータを取得する箇所を切り出したい
	// DB から一致する data を取得
	rows, err := db.Query("SELECT * FROM todos WHERE user_id = 'testUser'")

	if err != nil {
		log.Println(err)
	}

	var data []Todos

	// 1行ごとTODOにEntityをマッピングし、返却用のスライスに追加
	for rows.Next() {
		todo := Todos{}
		err = rows.Scan(&todo.todoId, &todo.title, &todo.userId, &todo.created_at, &todo.updated_at)
		if err != nil {
			log.Print(err)
			return
		}
		data = append(data, todo)
	}

	// DB から取得した data を Json 構造体 にマッピングする
	var todoResponses []EncodeTodo
	for _, v := range data {
		todoResponses = append(todoResponses, EncodeTodo{
			TodoId: v.todoId,
			Title:  v.title,
			UserId: v.userId,
		})
	}
	output, _ := json.MarshalIndent(todoResponses, "", "\t\t")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(output)
}

// AddTodo クライアントから送られてきたデータをもとに DB に追加
func AddTodo(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var err error
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	var todoRequest TodoRequest

	err = json.Unmarshal(body, &todoRequest)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Connection Failed"))
		return
	}

	todo := Todos{
		todoId: strconv.Itoa(rand.Int()) + "_todo",
		title:  todoRequest.Title,
		userId: todoRequest.UserId,
	}

	_, err = db.Exec("INSERT INTO todos (todo_id , title , user_id) VALUES ($1, $2 ,$3)", todo.todoId, todo.title, todo.userId)
	if err != nil {
		log.Print(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("success"))

}

// EditTodo クライアントから送られてきたデータをもとに DB を更新する
func EditTodo(db *sql.DB) {
	fmt.Println("PUT", db)
}

// DeleteTodo 指定データを DB から削除する
func DeleteTodo(db *sql.DB) {
	fmt.Println("DELETE", db)
}
