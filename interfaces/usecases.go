package interfaces

import (
	"github.com/postech-soat2-grupo16/pedidos-api/entities"
)

type OrderUseCase interface {
	List(clientID, status string) (*[]entities.Order, error)
	Create(order *entities.Order) (*entities.Order, error)
	GetByID(orderID string) (*entities.Order, error)
	Update(orderID string, updatedOrder *entities.Order) (*entities.Order, error)
	UpdateOrderStatus(orderID string, orderStatus entities.Status) (*entities.Order, error)
	Delete(orderID string) error
}
