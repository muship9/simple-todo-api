package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("simple-todo-api")
	server := http.Server{
		Addr: ":8080",
	}
	server.ListenAndServe()
}
