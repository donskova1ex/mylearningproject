package consumers

import (
	"github.com/IBM/sarama"
	"log"
)

// TODO: проверить
type RecipesConsumer struct {
	ready chan bool
}

func NewRecipesConsumer() *RecipesConsumer {
	return &RecipesConsumer{
		ready: make(chan bool),
	}
}

func (c *RecipesConsumer) Setup(sarama.ConsumerGroupSession) error {
	close(c.ready)
	return nil
}
func (c *RecipesConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *RecipesConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Println("Consumer kafka messages chan closed")
				return nil
			}
			log.Printf("Consumer kafka message: %s\n", message.Value)
			session.MarkMessage(message, "read")
		case <-session.Context().Done():
			return nil
		}
	}
}
