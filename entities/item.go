package entities

import (
	"time"
)

type Category string

type Item struct {
	ItemID      string    `json:"item_id"`
	Name        string    `json:"name"`
	Category    Category  `json:"category"`
	Description string    `json:"description;"`
	Price       float32   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
