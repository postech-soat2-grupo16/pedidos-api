package interfaces

import (
	"github.com/postech-soat2-grupo16/pedidos-api/entities"
)

type OrderGatewayI interface {
	Save(order *entities.Order) (*entities.Order, error)
	Update(orderID string, order *entities.Order) (*entities.Order, error)
	Delete(order *entities.Order) error
	GetByID(orderID string) (*entities.Order, error)
	GetAll() (*[]entities.Order, error)
	GetAllByClientID(clientID string) (*[]entities.Order, error)
}

type QueueGatewayI interface {
	SendMessage(order *entities.Order) (*entities.Order, error)
}
