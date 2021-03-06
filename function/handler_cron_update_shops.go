package primap

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/sue445/primap/config"
	"github.com/sue445/primap/db"
	"github.com/sue445/primap/prismdb"
	"github.com/sue445/primap/util"
	"golang.org/x/sync/errgroup"
	"time"
)

const (
	topicID = "sls-shop-save-topic"
)

// CronUpdateShops is called from cloud scheduler
func CronUpdateShops(ctx context.Context, m *pubsub.Message) error {
	cleanup, err := initFunction(ctx, 1.0)
	if err != nil {
		return errors.WithStack(err)
	}
	defer cleanup()

	span := sentry.StartSpan(ctx, "CronUpdateShops")
	defer span.Finish()

	err = getAndPublishShops(ctx, config.GetProjectID())

	if err != nil {
		handleError(err)
		return errors.WithStack(err)
	}

	return nil
}

func getAndPublishShops(ctx context.Context, projectID string) error {
	prismdbClient, err := prismdb.NewClient()
	if err != nil {
		return errors.WithStack(err)
	}

	start1 := time.Now()
	shops, err := prismdbClient.GetAllShops()
	if err != nil {
		return errors.WithStack(err)
	}
	duration1 := time.Now().Sub(start1)
	fmt.Printf("[DEBUG] prismdbClient.GetAllShops (%s)\n", duration1)
	fmt.Printf("[INFO][getAndPublishShops] fetched shops=%d\n", len(shops))

	start4 := time.Now()
	shops = config.AggregateShops(shops)
	duration4 := time.Now().Sub(start4)
	fmt.Printf("[DEBUG] config.AggregateShops (%s)\n", duration4)
	fmt.Printf("[INFO][getAndPublishShops] aggregated shops=%d\n", len(shops))

	pubsubClient, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return errors.WithStack(err)
	}

	start2 := time.Now()
	var eg errgroup.Group
	for _, shop := range shops {
		// c.f. https://golang.org/doc/faq#closures_and_goroutines
		shop := shop
		eg.Go(func() error {
			err := publishShop(ctx, pubsubClient, shop)
			if err != nil {
				return errors.WithStack(err)
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return errors.WithStack(err)
	}
	duration2 := time.Now().Sub(start2)
	fmt.Printf("[DEBUG] publishShop (%s)\n", duration2)

	fmt.Printf("[INFO][getAndPublishShops] Published shops=%d\n", len(shops))

	start3 := time.Now()
	err = deleteRemovedShops(ctx, projectID, shops)
	if err != nil {
		return errors.WithStack(err)
	}
	duration3 := time.Now().Sub(start3)
	fmt.Printf("[DEBUG] deleteRemovedShops (%s)\n", duration3)

	return nil
}

func publishShop(ctx context.Context, client *pubsub.Client, shop *prismdb.Shop) error {
	topic := client.Topic(topicID)

	data, err := json.Marshal(shop)

	if err != nil {
		return errors.WithStack(err)
	}

	_, err = topic.Publish(ctx, &pubsub.Message{Data: data}).Get(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func deleteRemovedShops(ctx context.Context, projectID string, newShops []*prismdb.Shop) error {
	var newShopNames []string
	for _, shop := range newShops {
		newShopNames = append(newShopNames, shop.Name)
	}

	dao := db.NewShopDao(ctx, projectID)
	dbShopNames, err := dao.GetAllIDs()
	if err != nil {
		return errors.WithStack(err)
	}

	removedShopNames := util.SubtractSlice(dbShopNames, newShopNames)

	var eg errgroup.Group
	for _, name := range removedShopNames {
		// c.f. https://golang.org/doc/faq#closures_and_goroutines
		name := name
		eg.Go(func() error {
			err := dao.DeleteShop(name)
			if err != nil {
				return errors.WithStack(err)
			}
			fmt.Printf("[INFO][deleteRemovedShops] Deleted shop=%s\n", name)
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return errors.WithStack(err)
	}

	fmt.Printf("[INFO][deleteRemovedShops] Deleted shops=%d\n", len(removedShopNames))

	return nil
}
