package hotels

import (
	"github.com/Shopify/sarama"
	"github.com/gorilla/mux"
	"github.com/polygens/consumer/config"
)

// App contains the objects for the service
type App struct {
	router        *mux.Router
	config        *config.Config
	consumerGroup sarama.ConsumerGroup
}

var app *App

// Init creates and starts the consumer
func Init(router *mux.Router, cfg *config.Config) {
	app = &App{router, cfg, nil}

	app.startKafka()
	app.setupRoutes()
}

// Close is used to handle a gracefull shutdown of the service
func Close() {
	app.consumerGroup.Close()
}
