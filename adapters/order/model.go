package Order

import (
	"github.com/postech-soat2-grupo16/pedidos-api/entities"
)

type Order struct {
	OrderID      string        `json:"order_id"`
	ClientID     string        `json:"client_id"`
	Status       string        `json:"status"`
	OrderedItems []OrderedItem `json:"ordered_items"`
	Notes        string        `json:"notes"`
	CreatedAt    string        `json:"created_at"`
	UpdatedAt    string        `json:"updated_at"`
}

type OrderedItem struct {
	ItemID   string  `json:"item_id"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func (o *Order) orderItemToEntity() (itemList []entities.OrderedItem) {
	for _, orderedItem := range o.OrderedItems {
		itemList = append(itemList, entities.OrderedItem{
			ItemID:   orderedItem.ItemID,
			Quantity: orderedItem.Quantity,
		})
	}
	return itemList
}

func (o *Order) ToUseCaseEntity() *entities.Order {
	return &entities.Order{
		OrderID:      o.OrderID,
		Status:       entities.Status(o.Status),
		OrderedItems: o.orderItemToEntity(),
		Notes:        o.Notes,
		ClientID:     o.ClientID,
		CreatedAt:    o.CreatedAt,
		UpdatedAt:    o.UpdatedAt,
	}
}

func orderItemFromEntity(orderedItems []entities.OrderedItem) (itemList []OrderedItem) {
	for _, orderedItem := range orderedItems {
		itemList = append(itemList, OrderedItem{
			ItemID:   orderedItem.ItemID,
			Quantity: orderedItem.Quantity,
		})
	}
	return itemList
}

func FromUseCaseEntity(order *entities.Order) *Order {
	return &Order{
		OrderID:      order.OrderID,
		Status:       string(order.Status),
		OrderedItems: orderItemFromEntity(order.OrderedItems),
		Notes:        order.Notes,
		ClientID:     order.ClientID,
		CreatedAt:    order.CreatedAt,
		UpdatedAt:    order.UpdatedAt,
	}
}
