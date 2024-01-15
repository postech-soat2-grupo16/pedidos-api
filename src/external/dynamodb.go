package external

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

func GetDynamoDbClient() *dynamodb.DynamoDB {
	// Specify the endpoint for DynamoDB Local
	endpoint := "http://localhost:8000"

	// Create a new AWS session
	sess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String(endpoint),
		Region:   aws.String("us-east-1"),
	})
	if err != nil {
		fmt.Println("Error creating session:", err)
		os.Exit(1)
	}

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	return svc
}
