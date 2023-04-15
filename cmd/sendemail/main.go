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

)



func sendEmail() error {

	return nil
}


type EmailMessage struct {
}


func handleRequest(msg EmailMessage) error {

	// ...
	return sendEmail()
}


func main() {
	lambda.Start(handleRequest)
}
