package main

import (
	"github.com/Trodyy/go-basics-project/pkg/server"
	"net/http"
)

func main() {
	s := server.NewHttpServer("localhost" , 8080)
	s.HandleFunc("/hello" , func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	s.Start()

}