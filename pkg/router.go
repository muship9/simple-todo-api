package pkg

import (
	"net/http"
)

// これは無意味な定義かもしれない
const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type Router interface {
	HandleTodosRequest(w http.ResponseWriter, r *http.Request)
}

func HandleTodosRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case GET:
		GetTodos()
	case POST:
		PostTodo()
	case PUT:
		PutTodo()
	case DELETE:
		DeleteTodo()
	default:
		// 指定メソッド以外はアクションを実行しない
		w.WriteHeader(405)
	}
}
