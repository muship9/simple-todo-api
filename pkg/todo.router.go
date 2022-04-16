package pkg

import (
	"fmt"
)

// GetTodos DB からデータを取得して一覧を返す
func GetTodos() {
	fmt.Println("GET")
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
