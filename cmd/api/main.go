package main

import (
	"fmt"
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
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "testdb"),
	)

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

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
