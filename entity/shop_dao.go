package entity

import (
	"cloud.google.com/go/firestore"
	"context"
)

const (
	// c.f. https://firebase.google.com/docs/firestore/manage-data/transactions#batched-writes
	maxBatchCount = 500
)

// ShopDao represents a shop DAO for Firestore
type ShopDao struct {
	projectID string
}

// NewShopDao create a ShopDao instance
func NewShopDao(projectID string) *ShopDao {
	return &ShopDao{projectID: projectID}
}

// SaveShops save shops to Firestore
func (d *ShopDao) SaveShops(shops []*ShopEntity, revision string) error {
	if len(shops) > maxBatchCount {
		slicedShops := sliceShops(shops, maxBatchCount)

		for _, s := range slicedShops {
			err := d.SaveShops(s, revision)
			if err != nil {
				return err
			}
		}

		return nil
	}

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, d.projectID)

	if err != nil {
		return err
	}

	defer client.Close()

	batch := client.Batch()
	for _, shop := range shops {
		docRef := client.Collection(shopCollectionName).Doc(shop.Name)
		shop.Revision = revision
		batch.Set(docRef, shop.toFirestore())
	}

	_, err = batch.Commit(ctx)

	if err != nil {
		return err
	}

	return nil
}

func sliceShops(shops []*ShopEntity, n int) [][]*ShopEntity {
	if len(shops) <= n {
		return [][]*ShopEntity{shops}
	}

	var sliced [][]*ShopEntity
	for start := 0; start < len(shops); start += n {
		end := start + n
		if end < len(shops) {
			sliced = append(sliced, shops[start:end])
		} else {
			// Last slice
			sliced = append(sliced, shops[start:])
		}
	}

	return sliced
}

// GetShop returns shop Firestore
func (d *ShopDao) GetShop(name string) (*ShopEntity, error) {
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

	data := docsnap.Data()
	shop := fromFirestore(data)

	return shop, nil
}
