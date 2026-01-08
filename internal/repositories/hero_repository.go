package repositories

import (
	"github.com/Web-serfer/test/internal/models"
)

// HeroRepository предоставляет методы для работы с героями в хранилище
type HeroRepository struct {
	// В реальном приложении здесь будут поля для подключения к базе данных
	// например, db *sql.DB или другие зависимости
}

// NewHeroRepository создает новый экземпляр HeroRepository
func NewHeroRepository() *HeroRepository {
	return &HeroRepository{}
}

// GetHeroByID возвращает героя по ID из хранилища
func (repo *HeroRepository) GetHeroByID(id int) (*models.Hero, error) {
	// В реальном приложении здесь будет запрос к базе данных
	// Пока возвращаем заглушку
	hero := &models.Hero{
		ID:          id,
		Name:        "Stored Hero",
		Description: "A hero from repository",
		Power:       90,
	}
	return hero, nil
}

// GetAllHeroes возвращает все герои из хранилища
func (repo *HeroRepository) GetAllHeroes() ([]*models.Hero, error) {
	// В реальном приложении здесь будет запрос к базе данных
	heroes := []*models.Hero{
		{ID: 1, Name: "Stored Hero 1", Description: "A hero from repository", Power: 90},
		{ID: 2, Name: "Stored Hero 2", Description: "Another hero from repository", Power: 85},
		{ID: 3, Name: "Stored Hero 3", Description: "Yet another hero from repository", Power: 80},
	}
	return heroes, nil
}