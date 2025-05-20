package server

import (
	"fmt"
	"go1f/pkg/api"
	"net/http"
	"os"
)

func Run() {
	fmt.Println("Запуск сервера!")

	// Получаем порт из переменной окружения или используем порт по умолчанию
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = ":7540" // порт по умолчанию
	} else {
		// Добавляем двоеточие в начале, если его нет
		if port[0] != ':' {
			port = ":" + port
		}
	}

	api.Init()
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
