package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/DataDog/datadog-lambda-go/logs"
)


type EmailMessage struct {
}


func handler(ctx context.Context, msg EmailMessage) (string, error) {
    logs.Log("Handling request")

    // Your Lambda function logic here

    logs.Log(fmt.Sprintf("Request completed"))

    return "OK", nil
}


func main() {
	lambda.Start(handler)
}
