package main

import (
	"context"
	"log/slog"
	"os"
	"strings"

	"github.com/donskova1ex/mylearningproject/internal"
	"github.com/donskova1ex/mylearningproject/internal/consumers"
)

func main() {

	logJSONHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(logJSONHandler)
	logger.Info("application started")
	slog.SetDefault(logger)
	brokersEnv := os.Getenv("KAFKA_BROKERS")
	if brokersEnv == "" {
		logger.Error("kafka bokers is unset")
		os.Exit(1)
	}

	brokers := strings.Split(brokersEnv, ",")
	groupId := internal.KafkaRecipesConsumerGroup
	topic := internal.KafkaTopicCreateRecipes

	consumer := consumers.NewRecipesConsumer()
	consumerGroup := consumers.NewConsumerGroup(consumer, groupId, topic, brokers, logger)

	if err := consumerGroup.Run(context.Background()); err != nil {
		logger.Error("consumer group error", slog.String("err", err.Error()))
		os.Exit(1)
	}

}
