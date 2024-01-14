package entities

import (
	"golang.org/x/exp/slices"
	"time"
)

// Order Status
type Status string

const (
	CreatedOrdersStatus        Status = "CRIADO"
	ReceivedOrderStatus               = "RECEBIDO"
	CookingOrderStatus                = "EM_PREPARACAO"
	ReadyOrderStatus                  = "PRONTO"
	DeliveredOrderStatus              = "ENTREGUE"
	DoneOrderStatus                   = "FINALIZADO"
	ApprovedPaymentOrderStatus        = "APROVADO"
	DeclinedPaymentOrderStatus        = "NEGADO"
)

type Order struct {
	OrderID   string        `json:"order_id"`
	Items     []OrderedItem `json:"items"`
	Status    Status        `json:"status"`
	Notes     string        `json:"notes"`
	ClientID  string        `json:"client_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type OrderedItem struct {
	ItemID   string `json:"item_id"`
	Quantity int    `json:"quantity"`
}

func (p *Order) IsStatusValid() bool {
	status := []Status{CreatedOrdersStatus, ReceivedOrderStatus, CookingOrderStatus, ReadyOrderStatus, DeliveredOrderStatus, DoneOrderStatus}
	return slices.Contains(status, p.Status)
}
