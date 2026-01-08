package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/Web-serfer/test/internal/services"
)

// HeroHandler — это функция-обработчик для нашего маршрута.
func HeroHandler(w http.ResponseWriter, r *http.Request) {
	service := services.NewService()

	// Проверяем, есть ли параметр ID в URL
	id := 1 // по умолчанию
	if idParam := r.URL.Query().Get("id"); idParam != "" {
		if parsedID, err := strconv.Atoi(idParam); err == nil {
			id = parsedID
		}
	}

	hero, err := service.GetHero(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Устанавливаем заголовок, чтобы клиент знал, что мы отправляем JSON.
	json.NewEncoder(w).Encode(hero)                   // Кодируем и отправляем ответ.
}
