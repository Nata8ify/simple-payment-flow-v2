package main

import (
	"adapter-send/internal/model"
	m1155thb "adapter-send/internal/service/m1155thb"
	mthb "adapter-send/internal/service/mthb"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
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
	log.Println("Execute send MBTH Token")
	var payload model.Payload
	err := json.Unmarshal([]byte(content), &payload)
	if err != nil {
		panic(fmt.Sprintf("Unable to unmarshal payload: %v", err))
	}
	mthb.ExecuteSendMBTHToken(client, &payload, contractAddress)
}

func executeSendM1155BTHToken(client *ethclient.Client, content string, contractAddress string) {
	log.Println("Execute send M1155BTH Token")
	var payload model.Payload1155
	err := json.Unmarshal([]byte(content), &payload)
	if err != nil {
		panic(fmt.Sprintf("Unable to unmarshal payload: %v", err))
	}
	m1155thb.ExecuteSendM1155BTHBToken(client, &payload, contractAddress)
}

func executeReadM1155BTHToken(client *ethclient.Client, content string, contractAddress string) {
	log.Println("Execute read M1155BTH Token")
	var payload model.Payload1155
	err := json.Unmarshal([]byte(content), &payload)
	if err != nil {
		panic(fmt.Sprintf("Unable to unmarshal payload: %v", err))
	}
	m1155thb.ExecuteReadM1155BTHBToken(client, &payload, contractAddress)
}

func main() {
	_, isLambdaMode := os.LookupEnv("LAMBDA_TASK_ROOT")
	if isLambdaMode {
		lambda.Start(Handler)
	} else {
		exampleContentArgumentSendMTHB, exSendArgMTHBExist := os.LookupEnv("EXAMPLE_ARGUMENT_SEND_MTHB")
		exampleContentArgumentSendM1155THB, exSendArgM1155THBExist := os.LookupEnv("EXAMPLE_ARGUMENT_SEND_M1155THB")
		exampleContentArgumentReadM1155THB, exReadArgM1155THBExist := os.LookupEnv("EXAMPLE_ARGUMENT_READ_M1155THB")
		client, envContractAddress := initialClient()
		if exSendArgMTHBExist {
			executeSendMBTHToken(client, exampleContentArgumentSendMTHB, envContractAddress)
		} else if exSendArgM1155THBExist {
			executeSendM1155BTHToken(client, exampleContentArgumentSendM1155THB, envContractAddress)
		} else if exReadArgM1155THBExist {
			executeReadM1155BTHToken(client, exampleContentArgumentReadM1155THB, envContractAddress)
		}
	}
}
