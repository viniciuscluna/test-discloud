package httphandler

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/viniciuscluna/test-discloud/internal/core/domain"
	"github.com/viniciuscluna/test-discloud/internal/core/ports"
)

type CarHandler struct {
	service ports.CarService
}

func NewCarHandler(service ports.CarService) *CarHandler {
	return &CarHandler{service: service}
}

func (h *CarHandler) RegisterRoutes(app *fiber.App) {
	cars := app.Group("/cars")
	cars.Get("/", h.FindAll)
	cars.Get("/:id", h.FindByID)
	cars.Post("/", h.Create)
	cars.Put("/:id", h.Update)
	cars.Delete("/:id", h.Delete)
}

func (h *CarHandler) FindAll(c fiber.Ctx) error {
	cars, err := h.service.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(cars)
}

func (h *CarHandler) FindByID(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	car, err := h.service.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "car not found"})
	}
	return c.JSON(car)
}

func (h *CarHandler) Create(c fiber.Ctx) error {
	var car domain.Car
	if err := c.Bind().JSON(&car); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.service.Create(&car); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(car)
}

func (h *CarHandler) Update(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var car domain.Car
	if err := c.Bind().JSON(&car); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	car.ID = uint(id)
	if err := h.service.Update(&car); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(car)
}

func (h *CarHandler) Delete(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.service.Delete(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
