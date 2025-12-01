package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func handler(ctx context.Context, event MyEvent) (string, error) {
	fmt.Printf("ctx: %v\n", ctx)
	fmt.Printf("event: %v\n", event)

	return fmt.Sprintf("Hello, %s!", event.Name), nil
}

func main() {
	lambda.Start(handler)
}
