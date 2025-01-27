package consumers

import (
	"log"

	"github.com/IBM/sarama"
)

// TODO: проверить и задать вопросы
type RecipesConsumer struct {
}

func NewRecipesConsumer() *RecipesConsumer {
	return &RecipesConsumer{}
}
func (c *RecipesConsumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}
func (c *RecipesConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// TODO: прокинуть свой логгер
func (c *RecipesConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Println("Consumer kafka messages chan closed")
				return nil
			}
			log.Printf("Consumer kafka message: %s\n", string(message.Value))
			//saving message прокинуть процессор проверка на корректность обработки.
			//Будет пытаться вычитать до бесконечности, пока сообщение не будет вычитано
			//retry  при вычитвании сообщения.
			session.MarkMessage(message, "read")
			session.Commit()
		case <-session.Context().Done():
			return nil
		}
	}
}
