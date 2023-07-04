package primap

import (
	"cloud.google.com/go/functions/metadata"
	"context"
	"github.com/cockroachdb/errors"
	"github.com/getsentry/sentry-go"
	"github.com/sue445/primap/config"
	"log"
	"os"
	"time"
)

const (
	eventExpiration = 30 * time.Minute
)

// Cleanup should call with defer
type Cleanup func()

func initFunction(ctx context.Context, tracesSampleRate float64) (Cleanup, error) {
	projectID := os.Getenv("GCP_PROJECT")

	sentryDebug := os.Getenv("SENTRY_DEBUG") != ""
	sentryDsn := os.Getenv("SENTRY_DSN")

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              sentryDsn,
		Environment:      os.Getenv("SENTRY_ENVIRONMENT"),
		AttachStacktrace: true,
		Debug:            sentryDebug,
		Release:          os.Getenv("SENTRY_RELEASE"),
		TracesSampleRate: tracesSampleRate,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	config.Init(&config.InitParams{
		ProjectID: projectID,
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

func isExpiredEvent(ctx context.Context) (bool, error) {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		// Assume an error on the function invoker and try again.
		return false, errors.WithStack(err)
	}

	expiration := meta.Timestamp.Add(eventExpiration)
	currentTime := time.Now()

	if currentTime.After(expiration) {
		log.Printf("[INFO] Event is expired: eventTime=%v, currentTime=%v\n", meta.Timestamp, currentTime)

		sentry.ConfigureScope(func(scope *sentry.Scope) {
			scope.SetLevel(sentry.LevelInfo)
			scope.SetExtras(map[string]interface{}{
				"eventTime":   meta.Timestamp,
				"currentTime": currentTime,
			})
		})
		sentry.CaptureMessage("Event is expired")

		return true, nil
	}

	return false, nil
}
