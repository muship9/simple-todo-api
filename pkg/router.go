package pkg

import (
	"fmt"
	"net/http"
)

type Router interface {
	HandleTodosRequest(w http.ResponseWriter, r *http.Request)
}

func HandleTodosRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("GET")
	case "POST":
		fmt.Println("POST")
	case "PUT":
		fmt.Println("PUT")
	case "DELETE":
		fmt.Println("DELETE")
	}
}
