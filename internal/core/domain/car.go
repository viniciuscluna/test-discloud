package domain

import "time"

type Car struct {
	ID        uint      `json:"id"`
	Brand     string    `json:"brand"`
	Model     string    `json:"model"`
	Year      int       `json:"year"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
