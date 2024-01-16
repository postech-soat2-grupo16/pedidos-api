package order

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
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
		TableName:  "table_gsi",
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

	//Creating a DynamoDB Query Input Search by Key (order id)
	fetch := &dynamodb.QueryInput{
		TableName:              &g.TableName,
		KeyConditionExpression: aws.String("order_id = :order_id"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":order_id": {
				S: aws.String(orderID),
			},
		},
	}

	// Fetching the Order using query
	result, err := g.repository.Query(fetch)
	if err != nil {
		fmt.Printf("Error fetching order ID: %s\nerror: %s", orderID, err)
		return nil, err
	}

	if len(result.Items) == 0 {
		fmt.Printf("Order ID: %s does not exist", orderID)
		return nil, nil
	}

	// Unmarshalling the DynamoDB item into Orders
	var orders []entities.Order
	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &orders); err != nil {
		fmt.Printf("Error Unmarshalling order ID: %s\nerror: %s", orderID, err)
		return nil, err
	}

	return &orders[0], nil
}

func (g *Gateway) GetAll() (orders *[]entities.Order, err error) {

	// Scanning the table
	params := &dynamodb.ScanInput{
		TableName: &g.TableName,
	}

	// Perform Scan operation
	result, err := g.repository.Scan(params)
	if err != nil {
		fmt.Printf("Error scanning table %s - Error: %s", g.TableName, err)
		return
	}

	if len(result.Items) == 0 {
		fmt.Printf("Table %s is empty", g.TableName)
		return orders, nil
	}

	// Unmarshalling the DynamoDB item into Orders
	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &orders); err != nil {
		fmt.Printf("Error Unmarshalling table data: %s\nerror: %s", g.TableName, err)
		return nil, err
	}

	return orders, nil
}

func (g *Gateway) GetAllByClientID(clientID string) (orders *[]entities.Order, err error) {

	// Querying table by client_id - GSI
	query := &dynamodb.QueryInput{
		TableName:              &g.TableName,
		IndexName:              aws.String("ClientIdIndex"),
		KeyConditionExpression: aws.String("client_id = :client_id"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":client_id": {S: aws.String(clientID)},
		},
	}

	// Perform Query operation
	result, err := g.repository.Query(query)
	if err != nil {
		fmt.Printf("Error scanning table %s - Error: %s", g.TableName, err)
		return
	}

	if len(result.Items) == 0 {
		fmt.Printf("Table %s is empty", g.TableName)
		return orders, nil
	}

	// Unmarshalling the DynamoDB item into Orders
	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &orders); err != nil {
		fmt.Printf("Error Unmarshalling table data: %s\nerror: %s", g.TableName, err)
		return nil, err
	}

	return orders, nil
}
