package prishops

import (
	"github.com/pkg/errors"
	"github.com/sue445/primap/prismdb"
	"github.com/sue445/primap/util"
)

// GetAllShops get all shopList
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

		// Pripara is set up in the PrismStone Shop.
		// c.f. https://twitter.com/T_ARTS_PRETTY/status/1484043957709402115
		if util.Contains(shop.Series, "prismstone") {
			shop.Series = append(shop.Series, "pripara")
		}

		shops = append(shops, shop)
	}

	// append shopList to prismdb shops
	for _, shop := range shopList {
		shops = append(shops, shop)
	}

	return shops, nil
}
