package hotels

import (
	"context"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
)

func (app *App) startKafka() {
	config := sarama.NewConfig()
	config.ClientID = app.config.Kafka.ClientID

	var err error
	config.Version, err = sarama.ParseKafkaVersion(app.config.Kafka.Version)
	if err != nil {
		log.Fatalf("Invalid kafka version used: %s", err)
	}

	app.consumerGroup, err = sarama.NewConsumerGroup(app.config.Kafka.Brokers, app.config.Kafka.ClientID, config)
	if err != nil {
		log.Fatalf("Failed to create consumer group: %s", err)
	}

	go consume(app.consumerGroup, []string{app.config.Kafka.LocationInputTopic})
}

func consume(consumerGroup sarama.ConsumerGroup, topics []string) {
	for {
		err := consumerGroup.Consume(context.Background(), topics, &hotelConsumer{})
		if err != nil {
			log.Fatalf("Error from consumer: %s", err)
		}
	}
}
