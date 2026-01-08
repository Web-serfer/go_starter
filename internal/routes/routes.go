package routes

import (
	"net/http"
	"github.com/Web-serfer/test/internal/handlers"
)

// SetupRouter настраивает и возвращает новый маршрутизатор.
func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Регистрируем обработчик для пути "/", используя хендлер из пакета handlers.
	mux.HandleFunc("/", handlers.HeroHandler)

	return mux
}
