package gormadapter

import (
	"github.com/viniciuscluna/test-discloud/internal/core/domain"
	"github.com/viniciuscluna/test-discloud/internal/core/ports"
	"gorm.io/gorm"
)

type carRecord struct {
	gorm.Model
	Brand    string
	CarModel string `gorm:"column:model"`
	Year     int
	Color    string
}

func (carRecord) TableName() string {
	return "cars"
}

type carRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) ports.CarRepository {
	return &carRepository{db: db}
}

func toDomain(r *carRecord) *domain.Car {
	return &domain.Car{
		ID:        r.ID,
		Brand:     r.Brand,
		Model:     r.CarModel,
		Year:      r.Year,
		Color:     r.Color,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}

func toRecord(c *domain.Car) *carRecord {
	rec := &carRecord{
		Brand:    c.Brand,
		CarModel: c.Model,
		Year:     c.Year,
		Color:    c.Color,
	}
	if c.ID != 0 {
		rec.ID = c.ID
	}
	return rec
}

func (r *carRepository) Create(car *domain.Car) error {
	rec := toRecord(car)
	if err := r.db.Create(rec).Error; err != nil {
		return err
	}
	car.ID = rec.ID
	car.CreatedAt = rec.CreatedAt
	car.UpdatedAt = rec.UpdatedAt
	return nil
}

func (r *carRepository) FindByID(id uint) (*domain.Car, error) {
	var rec carRecord
	if err := r.db.First(&rec, id).Error; err != nil {
		return nil, err
	}
	return toDomain(&rec), nil
}

func (r *carRepository) FindAll() ([]domain.Car, error) {
	var records []carRecord
	if err := r.db.Find(&records).Error; err != nil {
		return nil, err
	}
	cars := make([]domain.Car, len(records))
	for i, rec := range records {
		cars[i] = *toDomain(&rec)
	}
	return cars, nil
}

func (r *carRepository) Update(car *domain.Car) error {
	rec := toRecord(car)
	return r.db.Save(rec).Error
}

func (r *carRepository) Delete(id uint) error {
	return r.db.Delete(&carRecord{}, id).Error
}
