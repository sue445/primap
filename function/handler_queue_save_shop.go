package primap

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/sue445/primap/config"
	"github.com/sue445/primap/db"
	"github.com/sue445/primap/prismdb"
)

// QueueSaveShop is called from pub/sub subscription
func QueueSaveShop(ctx context.Context, m *pubsub.Message) error {
	cleanup, err := initFunction(ctx)
	if err != nil {
		return err
	}
	defer cleanup()

	err = queueSaveShopHandler(ctx, m)

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

	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetTags(map[string]string{
			"shop.Name":       shop.Name,
			"shop.Prefecture": shop.Prefecture,
			"shop.Address":    shop.Address,
		})
		scope.SetExtras(map[string]interface{}{
			"shop.Series": shop.Series,
		})
	})

	return saveShop(ctx, config.GetProjectID(), &shop)
}

func saveShop(ctx context.Context, projectID string, shop *prismdb.Shop) error {
	dao := db.NewShopDao(ctx, projectID)

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
