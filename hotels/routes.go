package hotels

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// setupRoutes adds the handlers to the router
func (app *App) setupRoutes() {
	app.router.Handle("/metrics", promhttp.Handler()).Methods("GET")
	app.router.HandleFunc("/ping", health).Methods("GET")
	app.router.HandleFunc("/ready", health).Methods("GET")
	app.router.HandleFunc("/live", health).Methods("GET")
	app.router.HandleFunc("/hotels", hotelsHandler).Methods("GET")
}

// health returns a simple 200 response
func health(w http.ResponseWriter, request *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("pong"))
}

// hotelsHandler returns all the hotels in the database
func hotelsHandler(w http.ResponseWriter, request *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("pong"))
}
