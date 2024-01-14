package interfaces

import (
	"github.com/postech-soat2-grupo16/pedidos-api/entities"
)

type OrderGatewayI interface {
	Save(order entities.Order) (*entities.Order, error)
	Update(orderID string, pedido entities.Order) (*entities.Order, error)
	Delete(orderID string) error
	GetByID(orderID string) (*entities.Order, error)
	GetAll(conds ...interface{}) ([]entities.Order, error)
}
