package config

import (
	"github.com/sue445/primap/prismdb"
	"sort"
	"strings"
)

var similarShopNames map[string][]string

func init() {
	similarShopNames = map[string][]string{
		"ＳＯＹＵＧａｍｅＦｉｅｌｄ湘南": {
			"ＳＯＹＵ　Ｇａｍｅ　Ｆｉｅｌｄ湘南店",
		},
		"ＳＯＹＵ　ＺＡＵＲＵＳ　ＷＯＲＬＤ大森": {
			"ＳＯＹＵＺＡＵＲＵＳＷＯＲＬＤ大森店",
		},
		"そごう千葉": {
			"そごう千葉本館６階玩具売場",
		},
		"タイトーステーション　ＢＩＧＦＵＮ平和島": {
			"タイトーステーション BIGFUN平和島",
		},
		"プリズムストーン札幌": {
			"プリズムストーン  新札幌",
		},
		"モーリーファンタジー向ヶ丘": {
			"モーリーファンタジー向ケ丘",
		},
	}
}

// AggregateShops returns aggregated shops with similar name
func AggregateShops(shops []*prismdb.Shop) []*prismdb.Shop {
	var reversedSimilarShopNames = map[string]string{}

	for key, values := range similarShopNames {
		for _, value := range values {
			reversedSimilarShopNames[value] = key
		}
	}

	aggregatedShopsMap := map[string]*prismdb.Shop{}

	for _, shop := range shops {
		shopName := strings.TrimSuffix(shop.Name, "店")

		if reversedSimilarShopNames[shop.Name] != "" {
			shopName = reversedSimilarShopNames[shop.Name]
		}

		if aggregatedShopsMap[shopName] == nil {
			aggregatedShopsMap[shopName] = &prismdb.Shop{
				Name:       shopName,
				Address:    shop.Address,
				Prefecture: shop.Prefecture,
				Series:     shop.Series,
			}
		} else {
			// merge series
			for _, series := range shop.Series {
				aggregatedShopsMap[shopName].Series = append(aggregatedShopsMap[shopName].Series, series)
			}
		}
	}

	var sortedAggregatedShopNames []string
	for k := range aggregatedShopsMap {
		sortedAggregatedShopNames = append(sortedAggregatedShopNames, k)
	}
	sort.Strings(sortedAggregatedShopNames)

	var aggregatedShops []*prismdb.Shop

	for _, shopName := range sortedAggregatedShopNames {
		shop := aggregatedShopsMap[shopName]
		sort.Strings(shop.Series)
		aggregatedShops = append(aggregatedShops, shop)
	}

	return aggregatedShops
}
