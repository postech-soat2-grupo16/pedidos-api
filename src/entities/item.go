package entities

import (
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

type Category string

const (
	drink    Category = "BEBIDA"
	sandwich          = "LANCHE"
	dessert           = "SOBREMESA"
	sideDish          = "ACOMPANHAMENTO"
)

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

func (i *Item) IsCategoryValid() bool {
	categories := []Category{drink, sandwich, dessert, sideDish}
	return slices.Contains(categories, i.Category)
}

func (i *Item) IsNameNull() bool {
	return len(strings.TrimSpace(i.Name)) == 0
}

func (i *Item) IsPriceValid() bool {
	return i.Price >= 0
}

func (i *Item) CopyItemWithNewValues(name, description string, price float32, category Category) Item {
	return Item{
		ItemID:      i.ItemID,
		Name:        name,
		Category:    category,
		Description: description,
		Price:       price,
	}
}
