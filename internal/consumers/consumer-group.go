package consumers

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/donskova1ex/mylearningproject/internal"
	"log"
	"os"
	"strings"
)

// TODO: доделать
var (
	version = sarama.DefaultVersion
	oldest  = true
	verbose = true
)

func ConsumerGroupConnection() {
	log.Println("Starting a new Sarama group connection")

	if verbose {
		sarama.Logger = log.New(os.Stdout, "[Sarama] ", log.LstdFlags)
	}

	config := sarama.NewConfig()
	config.Version = version

	if oldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	consumer := RecipesConsumer{
		ready: make(chan bool),
	}

	ctx, cancel := context.WithCancel(context.Background())

	brokers := os.Getenv("KAFKA_BROKERS")
	group := internal.KafkaRecipesConsumerGroup

	client, err := sarama.NewConsumerGroup(strings.Split(brokers, ","), group, config)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

}
