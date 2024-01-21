package message

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/postech-soat2-grupo16/pedidos-api/entities"
	"os"
)

type Gateway struct {
	queueURL string
	queue    *sqs.SQS
}

func NewGateway(queueClient *sqs.SQS) *Gateway {
	return &Gateway{
		queueURL: os.Getenv("QUEUE_URL"),
		queue:    queueClient,
	}
}

func (g *Gateway) SendMessage(order *entities.Order) (*entities.Order, error) {
	// Convert the struct to a JSON string
	jsonString, err := json.Marshal(order)
	if err != nil {
		fmt.Println("Error parsing order to json string:", err)
		return nil, err
	}
	stringMessage := string(jsonString)

	//Build message
	message := &sqs.SendMessageInput{
		QueueUrl:    &g.queueURL,
		MessageBody: &stringMessage,
	}

	g.queue.SendMessage(message)

	return order, nil
}
