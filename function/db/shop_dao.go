package db

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/cockroachdb/errors"
	"google.golang.org/api/iterator"
	"sort"
	"time"
)

// ShopDao represents a shop DAO for Firestore
type ShopDao struct {
	projectID string
	ctx       context.Context
}

// NewShopDao create a ShopDao instance
func NewShopDao(ctx context.Context, projectID string) *ShopDao {
	return &ShopDao{ctx: ctx, projectID: projectID}
}

// SaveShop save shop to Firestore
func (d *ShopDao) SaveShop(shop *ShopEntity) error {
	client, err := firestore.NewClient(d.ctx, d.projectID)

	if err != nil {
		return errors.WithStack(err)
	}

	defer client.Close()

	shop.UpdatedAt = time.Now()

	docRef := client.Collection(shopCollectionName).Doc(shop.Name)
	_, err = docRef.Set(d.ctx, shop)

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// LoadShop returns shop Firestore
func (d *ShopDao) LoadShop(name string) (*ShopEntity, error) {
	client, err := firestore.NewClient(d.ctx, d.projectID)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer client.Close()

	docsnap, err := client.Collection(shopCollectionName).Doc(name).Get(d.ctx)

	if err != nil {
		if docsnap != nil {
			// Key is not found in firestore
			return nil, nil
		}

		return nil, errors.WithStack(err)
	}

	var shop ShopEntity
	err = docsnap.DataTo(&shop)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &shop, nil
}

// LoadOrCreateShop returns shop Firestore. If not found, create Shop
func (d *ShopDao) LoadOrCreateShop(name string) (*ShopEntity, error) {
	foundShop, err := d.LoadShop(name)
	if err != nil {
		return nil, errors.WithStack(err)
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
	client, err := firestore.NewClient(d.ctx, d.projectID)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer client.Close()

	itr := client.Collection(shopCollectionName).Where("deleted", "==", false).Documents(d.ctx)
	defer itr.Stop()

	var ids []string

	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return []string{}, errors.WithStack(err)
		}

		ids = append(ids, doc.Ref.ID)
	}

	sort.Strings(ids)
	return ids, nil
}

// DeleteShop delete shop from firestore
func (d *ShopDao) DeleteShop(name string) error {
	shop, err := d.LoadShop(name)
	if err != nil {
		return errors.WithStack(err)
	}

	shop.Deleted = true
	err = d.SaveShop(shop)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
