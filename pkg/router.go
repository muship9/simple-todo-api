package pkg

import (
	"database/sql"
	"log"
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
	db, err := sql.Open("postgres", "host=localhost user=admin password=admin dbname=simpleTodoDb sslmode=disable")
	log.Print(&r)
	if err != nil {
		log.Fatal(err)
	}
	switch r.Method {
	case GET:
		GetTodos(db)
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
