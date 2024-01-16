package entities

// Status Order
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
