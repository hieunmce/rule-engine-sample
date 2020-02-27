package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hieunmce/rule-engine-sample/handler"
)

func main() {
	lambda.Start(handler.HandleEnqueue)
}
