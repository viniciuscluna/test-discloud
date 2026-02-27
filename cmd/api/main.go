package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	carrepo "github.com/viniciuscluna/test-discloud/internal/adapters/driven/gorm"
	httphandler "github.com/viniciuscluna/test-discloud/internal/adapters/driving/http"
	"github.com/viniciuscluna/test-discloud/internal/application"
	pgdriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type carMigration struct {
	gorm.Model
	Brand    string
	CarModel string `gorm:"column:model"`
	Year     int
	Color    string
}

func (carMigration) TableName() string {
	return "cars"
}

func main() {
	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(pgdriver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&carMigration{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	repo := carrepo.NewCarRepository(db)
	service := application.NewCarService(repo)
	handler := httphandler.NewCarHandler(service)

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Car CRUD API - Hexagonal Architecture")
	})

	handler.RegisterRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
