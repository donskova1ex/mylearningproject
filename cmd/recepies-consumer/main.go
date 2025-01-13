package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"os"
	"strings"
)

func main() {
	config := sarama.NewConfig()

	config.Consumer.Return.Errors = true
	brokers := os.Getenv("KAFKA_BROKERS")
	consumer, err := sarama.NewConsumer(strings.Split(brokers, ","), config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()
	partitionConsumer, err := consumer.ConsumePartition("internal.KafkaTopicCreateRecipes", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()
	messages := partitionConsumer.Messages()
	for {
		select {
		case msg := <-messages:
			fmt.Printf("Message claimed :%s \n", string(msg.Value))
		}
	}

}
