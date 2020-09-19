package cron

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/sue445/primap/config"
	"github.com/sue445/primap/prismdb"
	"log"
	"net/http"
)

const (
	topicID = "shop-save-topic"
)

// UpdateShopsHandler returns handler of /cron/update_shops
func UpdateShopsHandler(w http.ResponseWriter, r *http.Request) {
	err := getAndPublishShops()

	if err != nil {
		handleError(w, err)
		return
	}

	fmt.Fprint(w, "ok")
}

func handleError(w http.ResponseWriter, err error) {
	log.Printf("[ERROR] %+v", err)
	sentry.CaptureException(err)
	w.WriteHeader(500)
	fmt.Fprint(w, err)
}

func getAndPublishShops() error {
	prismdbClient, err := prismdb.NewClient()
	if err != nil {
		return err
	}

	shops, err := prismdbClient.GetAllShops()
	if err != nil {
		return err
	}

	ctx := context.Background()
	pubsubClient, err := pubsub.NewClient(ctx, config.GetProjectID())
	if err != nil {
		return err
	}

	for _, shop := range shops {
		err := publishShop(ctx, pubsubClient, shop)
		if err != nil {
			return err
		}
	}

	return nil
}

func publishShop(ctx context.Context, client *pubsub.Client, shop *prismdb.Shop) error {
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
