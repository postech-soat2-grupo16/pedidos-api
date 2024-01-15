package order

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/postech-soat2-grupo16/pedidos-api/entities"
)

type Gateway struct {
	TableName  string
	repository *dynamodb.DynamoDB
}

func NewGateway(repository *dynamodb.DynamoDB) *Gateway {
	return &Gateway{
		TableName:  "order_table_2",
		repository: repository,
	}
}

func (g *Gateway) Save(order *entities.Order) (*entities.Order, error) {

	//Marshaling order to a DynamoDB MAP
	item, err := dynamodbattribute.MarshalMap(order)
	if err != nil {
		fmt.Println("Error marshaling to DynamoDB attribute map:", err)
		return nil, err
	}

	//Creating a DynamoDB Input Item
	input := &dynamodb.PutItemInput{
		TableName: &g.TableName,
		Item:      item,
	}

	//Saving Input Item
	_, err = g.repository.PutItem(input)
	if err != nil {
		fmt.Println("Error inserting item:", err)
		return nil, err
	}

	fmt.Println("Item inserted successfully")
	return order, nil
}

func (g *Gateway) Update(orderID string, order *entities.Order) (*entities.Order, error) {
	return nil, nil
}

func (g *Gateway) Delete(orderID string) error {
	return nil
}

func (g *Gateway) GetByID(orderID string) (*entities.Order, error) {
	return nil, nil
}

func (g *Gateway) GetAll(conds ...interface{}) (orders []entities.Order, err error) {
	return nil, nil
}
