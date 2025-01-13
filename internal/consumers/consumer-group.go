package consumers

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
)

// TODO: разобрать
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
	//config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	client, err := sarama.NewConsumerGroup(cg.brokers, cg.group, config)
	if err != nil {
		return fmt.Errorf("error creating consumer group client: %w", err) //TODO: обернуть в ошибку нормальную
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := client.Consume(ctx, []string{cg.topic}, cg.consumer); err != nil { //
				cg.logger.Error("error consuming messages", slog.String("err", err.Error()))
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()
	log.Println("Consumer up and ready") //TODO:поменять на свой логгер

	for keepRunning {
		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-sigterm:
			log.Println("terminating: via signal") //TODO:свой логгер
			keepRunning = false
		case <-ctx.Done():
			log.Println("terminating: context cancelled") //TODO:свой логгер
			keepRunning = false
		}
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		log.Panicf("Error cloasing client: %v", err) //TODO: свой логгер, уровень "err"
	}
	return nil

}
