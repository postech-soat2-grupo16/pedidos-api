package external

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetDynamoDbClient() *dynamodb.DynamoDB {
	// Create a new AWS session
	if os.Getenv("IS_LOCAL") == "true" {
		sess, err := session.NewSession(&aws.Config{
			Region:   aws.String("us-east-1"),
			Endpoint: aws.String("http://localhost:9000"),
		})
		if err != nil {
			fmt.Println("Error creating session:", err)
			os.Exit(1)
		}

		// Create DynamoDB client
		svc := dynamodb.New(sess)
		_, err = svc.CreateTable(&dynamodb.CreateTableInput{
			TableName: aws.String("orders"),
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String("order_id"),
					AttributeType: aws.String("S"),
				},
			},
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("order_id"),
					KeyType:       aws.String("HASH"),
				},
			},
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(5),
				WriteCapacityUnits: aws.Int64(5),
			},
		})

		if err != nil {
			fmt.Println("Got error calling CreateTable:")
			fmt.Println(err.Error())
		}

		return svc
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		fmt.Println("Error creating session:", err)
		os.Exit(1)
	}

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	fmt.Printf("DynamoDB Connected: %v\n", svc)
	return svc
}
