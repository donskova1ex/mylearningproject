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
	"github.com/donskova1ex/mylearningproject/internal"
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
		return fmt.Errorf("error creating consumer group client: %w", internal.ErrCreateConsumerGroup) //TODO: обернуть в ошибку нормальную
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
	//log.Println("Consumer up and ready") //TODO:поменять на свой логгер
	cg.logger.Info("Consumer up and ready", slog.String("info", "Consumer up and ready")) //TODO:поменять на свой логгер

	for keepRunning {
		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-sigterm:
			cg.logger.Info("terminating: via signal", slog.String("info", "terminating: via signal"))
			//log.Println("terminating: via signal") //TODO:свой логгер
			keepRunning = false
		case <-ctx.Done():
			cg.logger.Info("terminating: context cancelled", slog.String("info", "terminating: context cancelled"))
			//log.Println("terminating: context cancelled") //TODO:свой логгер
			keepRunning = false
		}
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		cg.logger.Error("error cloasing client: %w", internal.ErrClosingCosumerGroupClient)
		//log.Panicf("Error cloasing client: %v", err) //TODO: свой логгер, уровень "err"
		return fmt.Errorf("error closing consumer group client: %w", internal.ErrClosingCosumerGroupClient)

	}
	return nil

}
