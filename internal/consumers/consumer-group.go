package consumers

import (
	"context"
	"errors"
	"github.com/IBM/sarama"
	"github.com/donskova1ex/mylearningproject/internal"
	"log"
	"os"
	"strings"
	"sync"
)

// TODO: разобрать

func main() {
	log.Println("Starting a new Sarama consumer")

	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	recipeConsumer := NewRecipesConsumer()
	group := internal.KafkaRecipesConsumerGroup
	brokers := os.Getenv("KAFKA_BROKERS")
	topics := internal.KafkaTopicCreateRecipes

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(strings.Split(brokers, ","), group, config)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := client.Consume(ctx, strings.Split(topics, ","), recipeConsumer); err != nil {
				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					return
				}
				log.Panicf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
			recipeConsumer.ready = make(chan bool)
		}
	}()
	<-recipeConsumer.ready
	log.Println("Sarama consumer up and running!...")
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}
