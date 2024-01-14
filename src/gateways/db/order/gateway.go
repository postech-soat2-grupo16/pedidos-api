package order

import (
	"log"

	"github.com/postech-soat2-grupo16/pedidos-api/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	repository *gorm.DB
}

func NewGateway(repository *gorm.DB) *Repository {
	return &Repository{repository: repository}
}

func (o *Repository) Save(order entities.Order) (*entities.Order, error) {
	result := o.repository.Create(&order)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &order, nil
}

func (o *Repository) Update(orderID string, order entities.Order) (*entities.Order, error) {
	order.OrderID = orderID
	for i := range order.Items {
		order.Items[i].ItemID = orderID
	}

	result := o.repository.Session(&gorm.Session{FullSaveAssociations: false}).Updates(&order)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return &order, nil
}

func (o *Repository) Delete(orderID string) error {
	order := entities.Order{
		OrderID: orderID,
	}
	result := o.repository.Preload("Items.ItemID").Delete(&order)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	return nil
}

func (o *Repository) GetByID(orderID string) (*entities.Order, error) {
	order := entities.Order{
		OrderID: orderID,
	}
	result := o.repository.Preload(clause.Associations).Preload("Items.ItemID").Preload("Pagamentos").First(&order)
	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

func (o *Repository) GetAll(conds ...interface{}) (orders []entities.Order, err error) {
	return orders, err
}
