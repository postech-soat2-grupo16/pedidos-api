package entities

import (
	"golang.org/x/exp/slices"
)

type Order struct {
	OrderID      string        `json:"order_id"`
	ClientID     string        `json:"client_id"`
	Status       Status        `json:"status"`
	OrderedItems []OrderedItem `json:"ordered_items"`
	Notes        string        `json:"notes"`
	CreatedAt    string        `json:"created_at"`
	UpdatedAt    string        `json:"updated_at"`
}

func (p *Order) IsStatusValid() bool {
	status := []Status{CreatedOrdersStatus, ReceivedOrderStatus, CookingOrderStatus, ReadyOrderStatus, DeliveredOrderStatus, DoneOrderStatus}
	return slices.Contains(status, p.Status)
}
