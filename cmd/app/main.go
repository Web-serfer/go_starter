package main

import (
	"log"
	"net/http"
	"github.com/Web-serfer/test/internal/routes"
	"github.com/Web-serfer/test/configs"
)

func main() {
	config := configs.LoadConfig()
	router := routes.SetupRouter()

	log.Printf("Конфигурация загружена. Порт: %s", config.ServerPort)
	log.Printf("Запуск сервера на http://localhost%s", config.ServerPort)
	log.Fatal(http.ListenAndServe(config.ServerPort, router))
}
