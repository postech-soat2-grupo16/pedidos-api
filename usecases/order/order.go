package order

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/postech-soat2-grupo16/pedidos-api/entities"
	"github.com/postech-soat2-grupo16/pedidos-api/interfaces"
	"github.com/postech-soat2-grupo16/pedidos-api/util"
	"time"
)

type UseCase struct {
	orderGateway interfaces.OrderGatewayI
	queueGateway interfaces.QueueGatewayI
}

func NewUseCase(orderGateway interfaces.OrderGatewayI, queueGateway interfaces.QueueGatewayI) UseCase {
	return UseCase{
		orderGateway: orderGateway,
		queueGateway: queueGateway,
	}
}

func (o UseCase) List(clientID, status string) (orders *[]entities.Order, err error) {

	if clientID == "" {
		orders, err = o.getAllOrders()
	} else {
		orders, err = o.getAllOrdersByClientID(clientID)
	}

	if status != "" {
		orders, err = o.filterOrdersByStatus(status, orders)
	}

	return orders, nil
}

func (o UseCase) getAllOrders() (orders *[]entities.Order, err error) {
	orders, err = o.orderGateway.GetAll()
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (o UseCase) getAllOrdersByClientID(clientID string) (orders *[]entities.Order, err error) {
	orders, err = o.orderGateway.GetAllByClientID(clientID)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (o UseCase) filterOrdersByStatus(status string, orders *[]entities.Order) (*[]entities.Order, error) {
	if orders != nil {
		var filteredOrders []entities.Order
		for _, order := range *orders {
			if string(order.Status) == status {
				filteredOrders = append(filteredOrders, order)
			}
		}
		return &filteredOrders, nil
	}
	return nil, nil
}

func (o UseCase) Create(order *entities.Order) (*entities.Order, error) {
	var now = time.Now().String()

	order.OrderID = uuid.New().String()
	order.CreatedAt = now
	order.UpdatedAt = ""
	var orderCreated, err = o.orderGateway.Save(order)
	if err != nil {
		return nil, err
	}

	return o.queueGateway.SendMessage(orderCreated)
}

func (o UseCase) GetByID(orderID string) (*entities.Order, error) {
	order, err := o.orderGateway.GetByID(orderID)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o UseCase) Update(orderID string, updatedOrder *entities.Order) (*entities.Order, error) {
	order, err := o.GetByID(orderID)
	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, nil
	}

	if !updatedOrder.IsStatusValid() {
		return nil, util.NewErrorDomain(fmt.Sprintf("Status %s is not valid", updatedOrder.Status))
	}

	var now = time.Now().String()
	order.OrderedItems = updatedOrder.OrderedItems
	order.Status = updatedOrder.Status
	order.Notes = updatedOrder.Notes
	order.UpdatedAt = now

	return o.orderGateway.Save(order)
}

func (o UseCase) UpdateOrderStatus(orderID string, orderStatus entities.Status) (*entities.Order, error) {
	order, err := o.GetByID(orderID)
	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, nil
	}

	order.Status = orderStatus
	if !order.IsStatusValid() {
		return order, util.NewErrorDomain(fmt.Sprintf("Status %s is not valid", orderStatus))
	}

	return o.orderGateway.Save(order)
}

func (o UseCase) Delete(orderID string) error {
	order, err := o.GetByID(orderID)
	if err != nil {
		return err
	}

	if order == nil {
		return util.NewErrorDomain(fmt.Sprintf("Order ID %s not found", orderID))
	}

	return o.orderGateway.Delete(order)
}
