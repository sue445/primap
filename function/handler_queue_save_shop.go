package primap

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/getsentry/sentry-go"
	"github.com/sue445/primap/config"
	"github.com/sue445/primap/db"
	"github.com/sue445/primap/prismdb"
	"log"
	"strings"
)

// QueueSaveShop is called from pub/sub subscription
func QueueSaveShop(ctx context.Context, m *pubsub.Message) error {
	cleanup, err := initFunction(ctx, 0.01)
	if err != nil {
		return errors.WithStack(err)
	}
	defer cleanup()

	err = queueSaveShopHandler(ctx, m)

	if err != nil {
		handleError(err)
		return errors.WithStack(err)
	}

	return nil
}

func queueSaveShopHandler(ctx context.Context, m *pubsub.Message) error {
	expired, err := isExpiredEvent(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	if expired {
		return nil
	}

	var shop prismdb.Shop
	err = json.Unmarshal(m.Data, &shop)

	if err != nil {
		return errors.WithStack(err)
	}

	span := sentry.StartSpan(ctx, "QueueSaveShop", sentry.WithTransactionName(fmt.Sprintf("QueueSaveShop: %s", shop.Name)))
	defer span.Finish()

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
	if !strings.HasPrefix(shop.Address, shop.Prefecture) {
		log.Printf("[WARN] Prefecture is mismatched with address: Prefecture=%s, Address=%s\n", shop.Prefecture, shop.Address)

		sentry.ConfigureScope(func(scope *sentry.Scope) {
			scope.SetLevel(sentry.LevelWarning)
		})
		sentry.CaptureMessage("Prefecture is mismatched with address")
	}

	dao := db.NewShopDao(ctx, projectID)

	entity, err := dao.LoadOrCreateShop(shop.Name)
	if err != nil {
		return errors.WithStack(err)
	}

	if !entity.IsUpdated(shop) {
		return nil
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
