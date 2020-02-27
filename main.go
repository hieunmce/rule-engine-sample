package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hieunmce/rule-engine-sample/core"
	"github.com/hieunmce/rule-engine-sample/handler"
)

func SQSHandler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		var pt core.Patient
		err := json.Unmarshal([]byte(message.Body), &pt)
		if err != nil {
			log.Println("failed to parse sqsEvent", err)
			continue
		}

		err = handler.HandleEnqueue(&pt)
		if err != nil {
			log.Println("failed to handle sqsEvent", err)
		}
	}

	return nil
}

func main() {
	lambda.Start(SQSHandler)
}
