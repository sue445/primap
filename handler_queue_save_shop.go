package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/sue445/primap/server/config"
	"github.com/sue445/primap/server/db"
	"github.com/sue445/primap/server/prismdb"
	"time"
)

// QueueSaveShop is called from pub/sub subscription
func QueueSaveShop(ctx context.Context, m *pubsub.Message) error {
	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(2 * time.Second)

	err := queueSaveShopHandler(ctx, m)

	if err != nil {
		handleError(err)
		return err
	}

	return nil
}

func queueSaveShopHandler(ctx context.Context, m *pubsub.Message) error {
	var shop prismdb.Shop
	err := json.Unmarshal(m.Data, &shop)

	if err != nil {
		return errors.WithStack(err)
	}

	return saveShop(ctx, config.GetProjectID(), &shop)
}

func saveShop(ctx context.Context, projectID string, shop *prismdb.Shop) error {
	dao := db.NewShopDao(projectID)

	entity, err := dao.LoadOrCreateShop(shop.Name)
	if err != nil {
		return errors.WithStack(err)
	}

	entity.Prefecture = shop.Prefecture
	entity.Series = shop.Series
	entity.Deleted = false

	err = entity.UpdateAddressWithGeography(ctx, shop.Address)
	if err != nil {
		return errors.WithStack(err)
	}

	err = dao.SaveShop(entity)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
