package Order

import (
	"github.com/postech-soat2-grupo16/pedidos-api/entities"
)

type Order struct {
	OrderID      string        `json:"order_id"`
	OrderedItems []OrderedItem `json:"ordered_items"`
	Notes        string        `json:"notes"`
	ClientID     string        `json:"client_id"`
	Status       string        `json:"status"`
}

type OrderedItem struct {
	ItemID   string `json:"item_id"`
	Quantity int    `json:"quantity"`
}

func (o *Order) OrderItemToEntity() (list []entities.OrderedItem) {
	for _, OrderedItem := range o.OrderedItems {
		list = append(list, entities.OrderedItem{
			ItemID:   OrderedItem.ItemID,
			Quantity: OrderedItem.Quantity,
		})
	}
	return list
}

func (o *Order) ToEntity() entities.Order {
	return entities.Order{
		OrderID:  o.OrderID,
		Items:    o.OrderItemToEntity(),
		Status:   entities.Status(o.Status),
		Notes:    o.Notes,
		ClientID: o.ClientID,
	}
}
