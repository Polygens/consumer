package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/polygens/consumer/config"
	"github.com/polygens/consumer/hotels"
)

var version string

func main() {
	log.Infof("Starting %s version: %s", filepath.Base(os.Args[0]), version)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}

	logLvl, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatalf("Failed to set log level: %s", err)
	}

	log.SetLevel(logLvl)

	r := mux.NewRouter()

	hotels.Init(r, cfg)
	defer hotels.Close()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.HTTPPort), r))
}
