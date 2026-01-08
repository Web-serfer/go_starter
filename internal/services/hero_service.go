package services

import (
	"github.com/Web-serfer/test/internal/models"
	"github.com/Web-serfer/test/internal/repositories"
)

// HeroServiceInterface определяет интерфейс для работы с героями
type HeroServiceInterface interface {
	GetHero(id int) (*models.Hero, error)
	GetAllHeroes() ([]*models.Hero, error)
}

// Service структура, содержащая зависимости сервиса
type Service struct {
	repo *repositories.HeroRepository
}

// NewService создает новый экземпляр Service
func NewService() *Service {
	return &Service{
		repo: repositories.NewHeroRepository(),
	}
}

// GetHero возвращает героя по ID
func (s *Service) GetHero(id int) (*models.Hero, error) {
	return s.repo.GetHeroByID(id)
}

// GetAllHeroes возвращает все доступные герои
func (s *Service) GetAllHeroes() ([]*models.Hero, error) {
	return s.repo.GetAllHeroes()
}