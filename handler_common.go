package main

import (
	"context"
	"github.com/getsentry/sentry-go"
	"github.com/sue445/gcp-secretmanagerenv"
	"github.com/sue445/primap/server/config"
	"log"
	"os"
)

func init() {
	projectID := os.Getenv("GCP_PROJECT")

	sentryDebug := os.Getenv("SENTRY_DEBUG") != ""
	secretmanager, err := secretmanagerenv.NewClient(context.Background(), projectID)
	if err != nil {
		panic(err)
	}

	sentryDsn, err := secretmanager.GetValueFromEnvOrSecretManager("SENTRY_DSN", false)
	if err != nil {
		panic(err)
	}

	googleMapsAPIKey, err := secretmanager.GetValueFromEnvOrSecretManager("GOOGLE_MAPS_API_KEY", false)
	if err != nil {
		panic(err)
	}

	err = sentry.Init(sentry.ClientOptions{
		Dsn:              sentryDsn,
		AttachStacktrace: true,
		Debug:            sentryDebug,
	})
	if err != nil {
		panic(err)
	}

	config.Init(&config.InitParams{
		ProjectID:        projectID,
		GoogleMapsAPIKey: googleMapsAPIKey,
	})
}

func handleError(err error) {
	log.Printf("[ERROR] %+v", err)
	sentry.CaptureException(err)
}
