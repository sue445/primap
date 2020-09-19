package cron

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/sue445/primap/prismdb"
	"log"
	"net/http"
)

const (
	topicID = "shop-save-topic"
)

// UpdateMapHandler returns handler of /cron/update_map
func UpdateMapHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Impl

	fmt.Fprint(w, "ok")
}

func handleError(w http.ResponseWriter, err error) {
	log.Printf("[ERROR] %+v", err)
	sentry.CaptureException(err)
	w.WriteHeader(500)
	fmt.Fprint(w, err)
}

func publishShop(client *pubsub.Client, ctx context.Context, shop *prismdb.Shop) error {
	topic := client.Topic(topicID)

	data, err := json.Marshal(shop)

	if err != nil {
		return err
	}

	_, err = topic.Publish(ctx, &pubsub.Message{Data: data}).Get(ctx)
	if err != nil {
		return err
	}

	return nil
}
