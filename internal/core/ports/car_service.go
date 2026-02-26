package ports

import "github.com/viniciuscluna/test-discloud/internal/core/domain"

type CarService interface {
	Create(car *domain.Car) error
	FindByID(id uint) (*domain.Car, error)
	FindAll() ([]domain.Car, error)
	Update(car *domain.Car) error
	Delete(id uint) error
}
