package prishops

import (
	"github.com/pkg/errors"
	"github.com/sue445/primap/prismdb"
)

// GetAllShops get all shops
func GetAllShops() ([]*Shop, error) {
	client, err := prismdb.NewClient()
	if err != nil {
		return []*Shop{}, errors.WithStack(err)
	}

	prismdbShops, err := client.GetAllShops()
	if err != nil {
		return []*Shop{}, errors.WithStack(err)
	}

	var shops []*Shop
	for _, prismdbShop := range prismdbShops {
		shop := &Shop{
			Name:       prismdbShop.Name,
			Prefecture: prismdbShop.Prefecture,
			Address:    prismdbShop.Address,
			Series:     prismdbShop.Series,
		}
		shops = append(shops, shop)
	}

	return shops, nil
}
