package db

import (
	"cloud.google.com/go/firestore"
	"context"
	"time"
)

// ShopDao represents a shop DAO for Firestore
type ShopDao struct {
	projectID string
}

// NewShopDao create a ShopDao instance
func NewShopDao(projectID string) *ShopDao {
	return &ShopDao{projectID: projectID}
}

// SaveShop save shop to Firestore
func (d *ShopDao) SaveShop(shop *ShopEntity) error {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, d.projectID)

	if err != nil {
		return err
	}

	defer client.Close()

	shop.UpdatedAt = time.Now()

	docRef := client.Collection(shopCollectionName).Doc(shop.Name)
	_, err = docRef.Set(ctx, shop)

	if err != nil {
		return err
	}

	return nil
}

// LoadShop returns shop Firestore
func (d *ShopDao) LoadShop(name string) (*ShopEntity, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, d.projectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()

	docsnap, err := client.Collection(shopCollectionName).Doc(name).Get(ctx)

	if err != nil {
		if docsnap != nil {
			// Key is not found in firestore
			return nil, nil
		}

		return nil, err
	}

	var shop ShopEntity
	err = docsnap.DataTo(&shop)
	if err != nil {
		return nil, err
	}

	return &shop, nil
}
