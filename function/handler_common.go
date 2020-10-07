package primap

import (
	"context"
	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/sue445/gcp-secretmanagerenv"
	"github.com/sue445/primap/config"
	"log"
	"os"
	"time"
)

// Cleanup should call with defer
type Cleanup func()

func initFunction(ctx context.Context) (Cleanup, error) {
	projectID := os.Getenv("GCP_PROJECT")

	sentryDebug := os.Getenv("SENTRY_DEBUG") != ""
	secretmanager, err := secretmanagerenv.NewClient(ctx, projectID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	sentryDsn, err := secretmanager.GetValueFromEnvOrSecretManager("SENTRY_DSN", false)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	googleMapsAPIKey, err := secretmanager.GetValueFromEnvOrSecretManager("GOOGLE_MAPS_API_KEY", false)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = sentry.Init(sentry.ClientOptions{
		Dsn:              sentryDsn,
		AttachStacktrace: true,
		Debug:            sentryDebug,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	config.Init(&config.InitParams{
		ProjectID:        projectID,
		GoogleMapsAPIKey: googleMapsAPIKey,
	})

	return func() {
		// Flush buffered events before the program terminates.
		// Set the timeout to the maximum duration the program can afford to wait.
		sentry.Flush(2 * time.Second)
	}, nil
}

func handleError(err error) {
	log.Printf("[ERROR] %+v", err)
	sentry.CaptureException(err)
}
