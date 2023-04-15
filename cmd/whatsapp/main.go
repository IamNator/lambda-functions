package main

import (
	"encoding/json"
	"log"
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"

)



type WhatsappMessage struct {
}


func handler(_ context.Context, msg WhatsappMessage) error {

	// ...
	return nil
}


func main() {
	lambda.Start(handler)
}
