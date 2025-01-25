package main

import (
	"context"
	"net/http"
	"github.com/Trodyy/go-basics-project/internal/db"
	"github.com/Trodyy/go-basics-project/internal/db/postgres"
	"github.com/Trodyy/go-basics-project/pkg/config"
	"github.com/Trodyy/go-basics-project/pkg/server"
)

func main() {
	cnf := config.LoadConfigOrPanic()
	pg , err := postgres.NewPostgres(cnf)
	if err != nil {
		panic(err)
	}
	if err := db.Migrate(context.Background(), pg); err != nil {
		panic(err)
	}
	s := server.NewHttpServer(cnf.Server)
	s.HandleFunc("/hello" , func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	s.Start()

}