package main

import (
	"adapter-send/internal/model"
	internal "adapter-send/internal/service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
)

func Handler(ctx context.Context, event events.SQSEvent) error {
	client, envContractAddress := initialClient()
	for _, message := range event.Records {
		executeSendMBTHToken(client, message.Body, envContractAddress)
	}
	return nil
}

func initialClient() (*ethclient.Client, string) {

	/* Initial configuration variable */
	envNodeUrl, existed := os.LookupEnv("NODE_URL")
	if !existed {
		panic("Environment variable [NODE_URL] is required")
	}

	envContractAddress, existed := os.LookupEnv("CONTRACT_ADDRESS")
	if !existed {
		panic("Environment variable [CONTRACT_ADDRESS] is required")
	}

	fmt.Printf("NODE_URL: %v \n", envNodeUrl)
	fmt.Printf("CONTRACT_ADDRESS: %v \n", envContractAddress)

	/* Dial */
	client, err := ethclient.Dial(envNodeUrl)
	if err != nil {
		panic("Unable to connect to Ethereum client")
	}
	return client, envContractAddress

}

func executeSendMBTHToken(client *ethclient.Client, content string, contractAddress string) {
	var payload model.Payload
	err := json.Unmarshal([]byte(content), &payload)
	if err != nil {
		panic(fmt.Sprintf("Unable to unmarshal payload: %v", err))
	}
	internal.ExecuteSendMBTHToken(client, &payload, contractAddress)
}

func main() {
	_, isLambdaMode := os.LookupEnv("LAMBDA_TASK_ROOT")
	if isLambdaMode {
		lambda.Start(Handler)
	} else {
		exampleContentArgument := os.Getenv("EXAMPLE_ARGUMENT")
		client, envContractAddress := initialClient()
		executeSendMBTHToken(client, exampleContentArgument, envContractAddress)
	}
}
