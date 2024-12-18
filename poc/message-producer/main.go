package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"log"
	"os"
	"strconv"
	"time"
)

// Require [CONFIG_ITERATION], [CONFIG_QUEUE_URL], [MODE] as an environment variable
func main() {
	envIteration, envIterationExist := os.LookupEnv("CONFIG_ITERATION")
	if !envIterationExist {
		log.Println("CONFIG_AMOUNT is required (should be a positive number)")
		return
	}
	amount, err := strconv.Atoi(envIteration)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Start produce messages...")
	mode := os.Getenv("MODE")
	switch mode {
	case "SEND":
		produce(amount)
	case "SEND_BATCH":
		produceBatch(amount, 10)
	case "PURGE":
		purge()
	}
}

func purge() {
	ctx := context.Background()
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}
	sqsClient := sqs.NewFromConfig(sdkConfig)
	var queueUrl = os.Getenv("CONFIG_QUEUE_URL")
	purgeQueueInput := sqs.PurgeQueueInput{
		QueueUrl: &queueUrl,
	}
	purgeQueueOutput, err := sqsClient.PurgeQueue(context.Background(), &purgeQueueInput)
	if err != nil {
		return
	}
	log.Println(purgeQueueOutput.ResultMetadata)
}

func produce(amount int) {

	ctx := context.Background()
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}
	sqsClient := sqs.NewFromConfig(sdkConfig)

	var message sqs.SendMessageInput
	var queueUrl = os.Getenv("CONFIG_QUEUE_URL")
	var groupId = "0"
	for i := 0; i < amount; i++ {
		//var messageDeduplicationId = fmt.Sprintf("%d:%d", time.Now().Nanosecond(), i)
		//body := "{\"amount\": 1000000000000000000, \"id\": 3, \"to\": \"0x12703f56e01bDD405B9209755161bF3cd797d73B\",\"from\" : \"0xFE3B557E8Fb62b89F4916B721be55cEb828dBd73\", \"issuerPrivateKey\" : \"8f2a55949038a9610f50fb23b5883af3b4ecb3c3bb792cbcefbd1542c692be63\"}"
		body := fmt.Sprintf("{\"amount\": 1, \"id\": 0, \"to\": \"0x12703f56e01bDD405B9209755161bF3cd797d73B\",\"from\" : \"0xFE3B557E8Fb62b89F4916B721be55cEb828dBd73\", \"issuerPrivateKey\" : \"8f2a55949038a9610f50fb23b5883af3b4ecb3c3bb792cbcefbd1542c692be63\", \"uqKey\": \"%v\"}", fmt.Sprintf("%d:%d", time.Now().UnixMilli(), i))
		message = sqs.SendMessageInput{
			MessageBody: &body,
			QueueUrl:    &queueUrl,
			//MessageDeduplicationId: &messageDeduplicationId,
			MessageGroupId: &groupId,
		}

		sendMessage, err := sqsClient.SendMessage(context.Background(), &message)
		if err != nil {
			log.Printf("Unable to send a message, %v \n", err)
		}
		log.Println(fmt.Printf("Message Sent, messageId: %v, body: %v \n", sendMessage.MessageId, body))

	}

}

func produceBatch(amount int, perBatchSize int) {
	ctx := context.Background()
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}
	sqsClient := sqs.NewFromConfig(sdkConfig)
	var queueUrl = os.Getenv("CONFIG_QUEUE_URL")
	var groupId = "0"
	for i := 0; i < amount; i++ {
		batch := make([]types.SendMessageBatchRequestEntry, 10)
		for j := 0; j < perBatchSize; j++ {
			uqKey := fmt.Sprintf("%d:%d:%d", time.Now().UnixMilli(), i, j)
			body := fmt.Sprintf("{\"amount\": 1, \"id\": 0, \"to\": \"0x12703f56e01bDD405B9209755161bF3cd797d73B\",\"from\" : \"0xFE3B557E8Fb62b89F4916B721be55cEb828dBd73\", \"issuerPrivateKey\" : \"8f2a55949038a9610f50fb23b5883af3b4ecb3c3bb792cbcefbd1542c692be63\", \"uqKey\": \"%v\"}", uqKey)
			batch[j] = types.SendMessageBatchRequestEntry{
				Id:             &uqKey,
				MessageBody:    &body,
				MessageGroupId: &groupId,
			}
		}
		batchInput := sqs.SendMessageBatchInput{
			Entries:  batch,
			QueueUrl: &queueUrl,
		}
		sendMessage, err := sqsClient.SendMessageBatch(ctx, &batchInput)
		if err != nil {
			log.Printf("Unable to send a message, %v \n", err)
		}
		log.Println(fmt.Printf("Message Sent, success: %v, failed: %v \n", sendMessage.Successful, sendMessage.Failed))

	}
}
