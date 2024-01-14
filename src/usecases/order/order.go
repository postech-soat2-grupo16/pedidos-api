package order

import (
	"errors"

	"github.com/postech-soat2-grupo16/pedidos-api/entities"
	"github.com/postech-soat2-grupo16/pedidos-api/interfaces"
	"gorm.io/gorm"
)

type UseCase struct {
	orderGateway interfaces.OrderGatewayI
}

func NewUseCase(orderGateway interfaces.OrderGatewayI) UseCase {
	return UseCase{orderGateway: orderGateway}
}

func (o UseCase) List(status string) (orders []entities.Order, err error) {
	if status != "" {
		order := entities.Order{
			Status: entities.Status(status),
		}
		return o.orderGateway.GetAll(order)
	}

	return o.orderGateway.GetAll()
}

func (o UseCase) Create(pedido entities.Order) (*entities.Order, error) {
	return o.orderGateway.Save(pedido)
}

func (o UseCase) GetByID(orderID string) (*entities.Order, error) {
	order, err := o.orderGateway.GetByID(orderID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return order, nil
}

func (o UseCase) Update(orderID string, order entities.Order) (*entities.Order, error) {
	if _, err := o.orderGateway.GetByID(orderID); errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return o.orderGateway.Update(orderID, order)
}

func (o UseCase) UpdateOrderStatus(orderID string, orderStatus string) (*entities.Order, error) {
	order, err := o.orderGateway.GetByID(orderID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	order.Status = entities.Status(orderStatus)
	return o.orderGateway.Update(orderID, *order)
}

func (o UseCase) Delete(orderID string) error {
	return o.orderGateway.Delete(orderID)
}
