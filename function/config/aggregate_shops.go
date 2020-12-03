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
		"イトーヨーカドーセブンホームセンター金町": {
			"イトーヨーカドー金町",
		},
		"ゲームエムエムランド": {
			"ゲーム・エムエムランド",
		},
		"そごう千葉": {
			"そごう千葉本館６階玩具売場",
		},
		"セガワールド勿来": {
			"セガ勿来",
		},
		"タイトーＦステーションイオンモール浜松市野": {
			"タイトーＦステーション浜松市野",
		},
		"タイトーＦステーションヨドバシ博多": {
			"タイトーステーションヨドバシ博多",
		},
		"タイトーステーション　ＢＩＧＦＵＮ平和島": {
			"タイトーステーション BIGFUN平和島",
		},
		"タイトーステーション小田原シティーモール": {
			"タイトーステーション小田原シティーモールクレッセ",
		},
		"ＮＩＣＯＰＡウイングタウン岡崎": {
			"ニコパウイングタウン岡崎",
		},
		"プリズムストーン札幌": {
			"プリズムストーン  新札幌",
		},
		"モーリーファンタジー向ヶ丘": {
			"モーリーファンタジー向ケ丘",
		},
		"ゆめパークゆめタウン徳山": {
			"ゆめパーク徳山",
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
		// Remove "店" that isn't "本店"
		// FIXME: I want to use `(?<!本)店$`, but Go regexp doesn't support negative look-ahead
		shopName := ""
		if strings.HasSuffix(shop.Name, "本店") {
			shopName = shop.Name
		} else {
			shopName = strings.TrimSuffix(shop.Name, "店")
		}

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
