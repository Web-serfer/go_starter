package models

// Hero представляет собой модель героя
type Hero struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Power       int    `json:"power"`
}

// HeroService предоставляет методы для работы с героями
type HeroService struct {
	// Здесь могут быть зависимости, например, репозиторий
	// repo *repository.HeroRepository
}

// NewHeroService создает новый экземпляр HeroService
func NewHeroService() *HeroService {
	return &HeroService{}
}

// GetHero возвращает героя по ID
func (hs *HeroService) GetHero(id int) (*Hero, error) {
	// В реальном приложении здесь будет логика получения героя из базы данных
	hero := &Hero{
		ID:          id,
		Name:        "Super Hero",
		Description: "A powerful hero",
		Power:       100,
	}
	return hero, nil
}

// GetAllHeroes возвращает все доступные герои
func (hs *HeroService) GetAllHeroes() ([]*Hero, error) {
	// В реальном приложении здесь будет логика получения героев из базы данных
	heroes := []*Hero{
		{ID: 1, Name: "Super Hero 1", Description: "A powerful hero", Power: 100},
		{ID: 2, Name: "Super Hero 2", Description: "Another powerful hero", Power: 95},
		{ID: 3, Name: "Super Hero 3", Description: "Yet another powerful hero", Power: 90},
	}
	return heroes, nil
}