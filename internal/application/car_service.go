package application

import (
	"github.com/viniciuscluna/test-discloud/internal/core/domain"
	"github.com/viniciuscluna/test-discloud/internal/core/ports"
)

type carService struct {
	repo ports.CarRepository
}

func NewCarService(repo ports.CarRepository) ports.CarService {
	return &carService{repo: repo}
}

func (s *carService) Create(car *domain.Car) error {
	return s.repo.Create(car)
}

func (s *carService) FindByID(id uint) (*domain.Car, error) {
	return s.repo.FindByID(id)
}

func (s *carService) FindAll() ([]domain.Car, error) {
	return s.repo.FindAll()
}

func (s *carService) Update(car *domain.Car) error {
	return s.repo.Update(car)
}

func (s *carService) Delete(id uint) error {
	return s.repo.Delete(id)
}
