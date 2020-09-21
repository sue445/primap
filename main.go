package main

import (
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gorilla/mux"
	"github.com/sue445/gcp-secretmanagerenv"
	"github.com/sue445/primap/config"
	"github.com/sue445/primap/job"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	projectID := os.Getenv("GCP_PROJECT")

	sentryDebug := os.Getenv("SENTRY_DEBUG") != ""
	secretmanager, err := secretmanagerenv.NewClient(context.Background(), projectID)
	if err != nil {
		log.Fatalf("secretmanagerenv.NewClient: %s", err)
	}

	sentryDsn, err := secretmanager.GetValueFromEnvOrSecretManager("SENTRY_DSN", false)
	if err != nil {
		log.Fatalf("secretmanager.GetValueFromEnvOrSecretManager: %s", err)
	}

	googleMapsAPIKey, err := secretmanager.GetValueFromEnvOrSecretManager("GOOGLE_MAPS_API_KEY", true)
	if err != nil {
		log.Fatalf("secretmanager.GetValueFromEnvOrSecretManager: %s", err)
	}

	err = sentry.Init(sentry.ClientOptions{
		Dsn:              sentryDsn,
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
		ProjectID:        projectID,
		GoogleMapsAPIKey: googleMapsAPIKey,
	})

	r := mux.NewRouter()
	r.HandleFunc("/job/cron/update_shops", job.CronUpdateShopsHandler).Methods("POST")
	r.HandleFunc("/job/queue/save_shop", job.QueueSaveShopHandler).Methods("POST")
	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
