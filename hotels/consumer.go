package hotels

import (
	log "github.com/sirupsen/logrus"

	"github.com/Shopify/sarama"
)

type hotelConsumer struct{}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (c *hotelConsumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (c *hotelConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (c *hotelConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Debugf("Message received: value = %s", string(message.Value))
		session.MarkMessage(message, "")
	}

	return nil
}
