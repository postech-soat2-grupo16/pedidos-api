package entities

type OrderedItem struct {
	ItemID   string `json:"item_id"`
	Quantity int    `json:"quantity"`
}
