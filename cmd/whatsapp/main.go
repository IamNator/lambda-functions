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



func sendWhatsappMessage() error {

	return nil
}


type WhatsappMessage struct {
}


func handleRequest(msg WhatsappMessage) error {

	// ...
	return sendWhatsappMessage()
}


func main() {
	lambda.Start(handleRequest)
}
