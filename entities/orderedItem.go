package entities

type OrderedItem struct {
	ItemID      string  `json:"item_id"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
}
