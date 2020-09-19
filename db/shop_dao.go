package db

import (
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/iterator"
	"sort"
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

// LoadOrCreateShop returns shop Firestore. If not found, create Shop
func (d *ShopDao) LoadOrCreateShop(name string) (*ShopEntity, error) {
	foundShop, err := d.LoadShop(name)
	if err != nil {
		return nil, err
	}

	if foundShop != nil {
		return foundShop, nil
	}

	createdShop := &ShopEntity{
		Name:      name,
		CreatedAt: time.Now(),
		Series:    []string{},
	}
	return createdShop, nil
}

// GetAllIDs returns all shop ids
func (d *ShopDao) GetAllIDs() ([]string, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, d.projectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()

	itr := client.Collection(shopCollectionName).Where("deleted", "==", false).Documents(ctx)
	defer itr.Stop()

	var ids []string

	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return []string{}, err
		}

		ids = append(ids, doc.Ref.ID)
	}

	sort.Strings(ids)
	return ids, nil
}
