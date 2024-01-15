package interfaces

import (
	"github.com/postech-soat2-grupo16/pedidos-api/entities"
)

type OrderUseCase interface {
	List(status string) ([]entities.Order, error)
	Create(order *entities.Order) (*entities.Order, error)
	GetByID(orderID string) (*entities.Order, error)
	Update(orderID string, order *entities.Order) (*entities.Order, error)
	UpdateOrderStatus(orderID string, orderStatus string) (*entities.Order, error)
	Delete(orderID string) error
}
