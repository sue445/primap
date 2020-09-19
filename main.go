package main

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gorilla/mux"
	"github.com/sue445/primap/config"
	"github.com/sue445/primap/cron"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	sentryDebug := os.Getenv("SENTRY_DEBUG") != ""

	err := sentry.Init(sentry.ClientOptions{
		AttachStacktrace: true,
		Debug:            sentryDebug,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(2 * time.Second)

	config.Init(&config.InitParams{
		ProjectID:        os.Getenv("GCP_PROJECT"),
		GoogleMapsAPIKey: os.Getenv("GOOGLE_MAPS_API_KEY"),
	})

	r := mux.NewRouter()
	r.HandleFunc("/cron/update_map", cron.UpdateMapHandler).Methods("POST")
	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
