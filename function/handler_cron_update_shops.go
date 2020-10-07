package primap

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sue445/primap/config"
	"github.com/sue445/primap/db"
	"github.com/sue445/primap/prismdb"
	"github.com/sue445/primap/util"
)

const (
	topicID = "sls-shop-save-topic"
)

// CronUpdateShops is called from cloud scheduler
func CronUpdateShops(ctx context.Context, m *pubsub.Message) error {
	cleanup, err := initFunction(ctx)
	if err != nil {
		return err
	}
	defer cleanup()

	err = getAndPublishShops(ctx, config.GetProjectID())

	if err != nil {
		handleError(err)
		return err
	}

	return nil
}

func getAndPublishShops(ctx context.Context, projectID string) error {
	prismdbClient, err := prismdb.NewClient()
	if err != nil {
		return errors.WithStack(err)
	}

	shops, err := prismdbClient.GetAllShops()
	if err != nil {
		return errors.WithStack(err)
	}

	pubsubClient, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return errors.WithStack(err)
	}

	for _, shop := range shops {
		err := publishShop(ctx, pubsubClient, shop)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	fmt.Printf("[INFO][getAndPublishShops] Published shops=%d\n", len(shops))

	err = deleteRemovedShops(ctx, projectID, shops)
	if err != nil {
		return errors.WithStack(err)
	}

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

	for _, name := range removedShopNames {
		err := dao.DeleteShop(name)
		if err != nil {
			return errors.WithStack(err)
		}
		fmt.Printf("[INFO][deleteRemovedShops] Deleted shop=%s\n", name)
	}

	fmt.Printf("[INFO][deleteRemovedShops] Deleted shops=%d\n", len(removedShopNames))

	return nil
}
