package main

import (
	"github.com/Trodyy/go-basics-project/pkg/server"
	"net/http"
	"github.com/Trodyy/go-basics-project/pkg/config"
)

func main() {
	cnf := config.LoadConfigOrPanic()
	s := server.NewHttpServer(cnf.Server)
	s.HandleFunc("/hello" , func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	s.Start()

}