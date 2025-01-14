package main

import (
	"log/slog"
	"os"

	"github.com/donskova1ex/mylearningproject/internal"
	"github.com/donskova1ex/mylearningproject/internal/consumers"
)

func main() {
	// TODO: прочитать энв и т.д., создать косьюмер, косьюмер группу, вызвать ран

	logJSONHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(logJSONHandler)
	logger.Info("application started")
	slog.SetDefault(logger)

	brokers := os.Getenv("KAFKA_BROKERS")
	groupId := internal.KafkaRecipesConsumerGroup
	recipeConsumer := consumers.NewRecipesConsumer()
}
