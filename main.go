package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"time"
)

// response struct that supports JSON serialization and
// defines the HTTP response body of a successful invocation of your AWS Lambda function.
type response struct {
	UTC time.Time `json:"utc"`
}

// handleRequest creates a response struct containing the current time in UTC and
// then proceeds to serialize it as JSON. In case the serialization fails,
// It return the error; if everything goes well, It respond with serialized JSON as the response body and a status code of 200.
func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	now := time.Now()
	resp := &response{
		UTC: now.UTC(),
	}

	body, err := json.Marshal(resp)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}

func main() {
	// Register handler function in the main function using the AWS Lambda for Go library.
	lambda.Start(handleRequest)
}
