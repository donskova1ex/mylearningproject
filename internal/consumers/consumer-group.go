package consumers

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
	"github.com/donskova1ex/mylearningproject/internal"
)

// TODO: разобрать

func main() {
	keepRunning := true
	recipeConsumer := NewRecipesConsumer()
	group := internal.KafkaRecipesConsumerGroup
	brokers := os.Getenv("KAFKA_BROKERS")
	topics := internal.KafkaTopicCreateRecipes

	config := sarama.NewConfig()
	config.Version = sarama.DefaultVersion
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

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
			if err := client.Consume(ctx, []string{topics}, recipeConsumer); err != nil { //
				log.Panicf("Error consuming messages: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
			recipeConsumer.ready = make(chan bool)
		}
	}()
	<-recipeConsumer.ready
	log.Println("Consumer up and ready")

	for keepRunning {
		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-sigterm:
			log.Println("terminating: via signal")
			keepRunning = false
		case <-ctx.Done():
			log.Println("terminating: context cancelled")
			keepRunning = false
		}
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		log.Panicf("Error cloasing client: %v", err)
	}

}
