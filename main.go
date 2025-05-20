package main

import (
	"github.com/Dowel/TODO_LIST_YANDEX/pkg/db"
	"go1f/pkg/server"
	"net/http"
	"os"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./web")))
	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "scheduler.db" // значение по умолчанию
	}

	db.Init(dbFile)
	defer db.Close()
	server.Run()
}
