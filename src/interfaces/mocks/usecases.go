package mocks

import (
	"github.com/postech-soat2-grupo16/pedidos-api/entities"
	"github.com/stretchr/testify/mock"
)

type OrderUseCase struct {
	mock.Mock
}

func (_m *OrderUseCase) List(clientID, status string) (*[]entities.Order, error) {
	ret := _m.Called(clientID, status)

	var r0 *[]entities.Order
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*[]entities.Order)
	}
	var r1 error = ret.Error(1)

	return r0, r1
}

func (_m *OrderUseCase) Create(order *entities.Order) (*entities.Order, error) {
	ret := _m.Called(order)

	var r0 *entities.Order
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*entities.Order)
	}
	var r1 error = ret.Error(1)

	return r0, r1
}

func (_m *OrderUseCase) GetByID(orderID string) (*entities.Order, error) {
	ret := _m.Called(orderID)

	var r0 *entities.Order
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*entities.Order)
	}
	var r1 error = ret.Error(1)

	return r0, r1
}

func (_m *OrderUseCase) Update(orderID string, updatedOrder *entities.Order) (*entities.Order, error) {
	ret := _m.Called(orderID, updatedOrder)

	var r0 *entities.Order
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*entities.Order)
	}
	var r1 error = ret.Error(1)

	return r0, r1
}

func (_m *OrderUseCase) UpdateOrderStatus(orderID string, orderStatus entities.Status) (*entities.Order, error) {
	ret := _m.Called(orderID, orderStatus)

	var r0 *entities.Order
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*entities.Order)
	}
	var r1 error = ret.Error(1)

	return r0, r1
}

func (_m *OrderUseCase) Delete(orderID string) error {
	ret := _m.Called(orderID)

	var r0 error = ret.Error(0)

	return r0
}
