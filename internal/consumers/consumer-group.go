package consumers

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
)

type ConsumerGroup struct {
	consumer sarama.ConsumerGroupHandler
	group    string
	brokers  []string
	topic    string
	logger   *slog.Logger
}

func NewConsumerGroup(
	consumer sarama.ConsumerGroupHandler,
	group string,
	topic string,
	brokers []string,
	logger *slog.Logger,
) *ConsumerGroup {
	return &ConsumerGroup{
		consumer: consumer,
		group:    group,
		topic:    topic,
		brokers:  brokers,
		logger:   logger,
	}
}

func (cg *ConsumerGroup) Run(ctx context.Context) error {

	keepRunning := true

	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	client, err := sarama.NewConsumerGroup(cg.brokers, cg.group, config)
	if err != nil {
		return fmt.Errorf("error creating consumer group client: %w", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := client.Consume(ctx, []string{cg.topic}, cg.consumer); err != nil { //
				cg.logger.Error("error consuming messages",
					slog.String("err", err.Error()))
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()
	cg.logger.Info("Consumer up and ready")

	for keepRunning {
		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-sigterm:
			cg.logger.Info("terminating: via signal")
			keepRunning = false
		case <-ctx.Done():
			cg.logger.Info("terminating: context cancelled")
			keepRunning = false
		}
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		return fmt.Errorf("error closing consumer group client: %w", err)
	}
	return nil

}
