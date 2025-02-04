package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/IBM/sarama"
	"github.com/donskova1ex/mylearningproject/internal"
	"github.com/donskova1ex/mylearningproject/internal/consumers"
	"github.com/donskova1ex/mylearningproject/internal/processors"
	"github.com/donskova1ex/mylearningproject/internal/repositories"
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

	pgDSN := os.Getenv("POSTGRES_DSN")
	if pgDSN == "" {
		logger.Error("empty POSTGRES_DSN")
		os.Exit(1)
	}

	db, err := repositories.NewPostgresDB(pgDSN)
	if err != nil {
		logger.Error("can not create postgres db connection", slog.String("error", err.Error()))
		return
	}
	defer db.Close()

	repository := repositories.NewRepository(db, logger)
	recipesProcessor := processors.NewRecipe(repository, logger)

	brokers := strings.Split(brokersEnv, ",")
	groupId := internal.KafkaRecipesConsumerGroup
	topic := internal.KafkaTopicCreateRecipes

	if err := initTopics(brokers, sarama.NewConfig()); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	consumer := consumers.NewRecipesConsumer(recipesProcessor)
	consumerGroup := consumers.NewConsumerGroup(consumer, groupId, topic, brokers, logger)

	if err := consumerGroup.Run(context.Background()); err != nil {
		logger.Error("consumer group error", slog.String("err", err.Error()))
		os.Exit(1)
	}

}

func initTopics(brokers []string, config *sarama.Config) error {
	admin, err := sarama.NewClusterAdmin(brokers, config)
	if err != nil {
		return fmt.Errorf("error while creating cluster admin: %w", err)
	}
	defer func() { _ = admin.Close() }()
	topics, err := admin.ListTopics()
	if err != nil {
		return fmt.Errorf("error while getting topics: %w", err)
	}
	if _, exists := topics[internal.KafkaTopicCreateRecipes]; !exists {
		err = admin.CreateTopic(internal.KafkaTopicCreateRecipes, &sarama.TopicDetail{
			NumPartitions:     1,
			ReplicationFactor: 3,
		}, false)
		if err != nil {
			return fmt.Errorf("error while creating topic: %w", err)
		}
	}

	return nil
}
