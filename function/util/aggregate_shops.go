package util

import (
	"github.com/sue445/primap/prismdb"
	"golang.org/x/text/width"
	"regexp"
	"sort"
	"strings"
)

var similarShopNames map[string][]string

func init() {
	similarShopNames = map[string][]string{
		"AMジャムジャムつくば": {
			"ジャムジャムつくば",
		},
		"ＮＩＣＯＰＡイオン伊勢": {
			"ＮＩＣＯＰＡ伊勢",
		},
		"ＮＩＣＯＰＡなるぱーく": {
			"にこぱなるぱーく",
		},
		"ＳＯＹＵ　Ｆａｍｉｌｙ　Ｇａｍｅ　Ｆｉｅｌｄ花巻": {
			"ＳＯＹＵファミリーゲームフィールド花巻",
		},
		"ＳＯＹＵＧａｍｅＦｉｅｌｄ湘南": {
			"ＳＯＹＵ　Ｇａｍｅ　Ｆｉｅｌｄ湘南店",
		},
		"ＳＯＹＵＴＯＹ’ｓＮＹ守谷": {
			"ソユートイズニューヨーク 守谷",
		},
		"アピタ名古屋北": {
			"アピタ名北",
		},
		"アピナ新利府北館": {
			"アピナ新利府 北館",
		},
		"アピタプラス岩倉": {
			"アピタ岩倉",
		},
		"イオンモール熱田": {
			"イオン熱田",
		},
		"イオンモールつくば": {
			"イオンつくば",
		},
		"イオン新百合ヶ丘ファミリーパーク": {
			"イオン新百合丘",
		},
		"イトーヨーカドーアリオ柏": {
			"イトーヨーカドーセブンパークアリオ柏",
		},
		"イトーヨーカドーセブンホームセンター金町": {
			"イトーヨーカドー金町",
		},
		"おもちゃのヨシダ本店": {
			"おもちゃのヨシダ",
		},
		"カーニバル・Ｃドーム": {
			"カーニバルシードーム",
		},
		"ゲームエムエムランド": {
			"ゲーム・エムエムランド",
		},
		"シルクハット川崎モアーズ": {
			"シルクハットモアーズ",
		},
		"そごう千葉": {
			"そごう千葉本館６階玩具売場",
		},
		"セガワールド勿来": {
			"セガ勿来",
		},
		"ソユーザウルスワールド大森": {
			"ＳＯＹＵＺＡＵＲＵＳＷＯＲＬＤ大森店",
			"ＳＯＹＵ　ＺＡＵＲＵＳ　ＷＯＲＬＤ大森",
		},
		"ソユーフォレストハンター松前": {
			"ＳＯＹＵＦｏｒｅｓｔＨｕｎｔｅｒ松前",
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
		"天王寺パスカ": {
			"天王子パスカ",
		},
		"トイザらス・ベビーザらス広島府中": {
			"トイザらスベビーザらス広島府中",
		},
		"博品館TOYPARK": {
			"博品館TOYPARK銀座本館",
		},
		"ハローズガーデン海老名": {
			"ハローガーデン海老名",
		},
		"バロー萩原": {
			"ハローズ萩原",
		},
		"ビックカメラ名古屋JRゲートタワー": {
			"ビックカメラＪＲゲートタワー",
		},
		"ビックカメラ有楽町": {
			"ビックカメラ有楽町本館",
		},
		"ビックカメラ水戸駅": {
			"ビックカメラ水戸",
			"ビックカメラ水戸駅前",
		},
		"ビックトイズプライムツリー赤池": {
			"ビックカメラビックカメラプライムツリー赤池",
		},
		"ビックロビックカメラ新宿東口": {
			"ビックカメラビックロ新宿東口",
		},
		"プラサカプコン府中": {
			"プラザカプコン府中",
		},
		"プリズムストーン札幌": {
			"プリズムストーン  新札幌",
		},
		"プリズムストーン仙台": {
			"アニメガ×ソフマップ仙台駅前",
		},
		"プリズムストーンカフェ原宿": {
			"プリズムストーン  原宿",
		},
		"プリズムストーン名古屋": {
			"アニメガ×ソフマップ名古屋",
		},
		"プリズムストーンなんば": {
			"アニメガ×ソフマップなんば",
		},
		"プリズムストーン福岡": {
			"アニメガ×ソフマップ天神1号館",
		},
		"プリズムストーン神戸": {
			"アニメガ×ソフマップ神戸ハーバーランド",
		},
		"水戸京成百貨": {
			"水戸京成百貨店７階玩具売場",
		},
		"モーリーファンタジー向ヶ丘": {
			"モーリーファンタジー向ケ丘",
		},
		"モーリーファンタジーメガドンキ大和": {
			"モーリーファンタジーピアゴ大和",
		},
		"モーリーファンタジーイオン相模原": {
			"モーリーファンタジーイオン相模",
		},
		"モーリーファンタジー千葉ニュータウン": {
			"モーリーファンタジー千葉ニュ-タウン",
		},
		"モーリーファンタジーf加西北条": {
			"モーリーファンタジー加西北条",
		},
		"モーリーファンタジー・ｆ綾川": {
			"モーリーファンタジーf綾川",
			"モーリーファンタジー綾川",
		},
		"モーリーファンタジー・ｆ岡山": {
			"モーリーファンタジーf岡山",
		},
		"モーリーファンタジー・ｆ新瑞橋": {
			"モーリーファンタジーf新瑞橋",
		},
		"モーリーファンタジー・ｆ広島祇園": {
			"モーリーファンタジーf広島祗園",
			"モーリーファンタジー広島祇園",
			"モーリーファンタジー広島祗園",
		},
		"ヤマダ電機LABI品川大井町": {
			"LABI 品川大井町 住まいる家電館",
		},
		"ゆめパークゆめタウン徳山": {
			"ゆめパーク徳山",
		},
		"ヨドバシカメラ新宿西口本店": {
			"ヨドバシカメラ新宿西口",
		},
		"ヨドバシカメラ千葉": {
			"ヨドバシカメラマルチメディア千葉",
		},
		"ヨドバシカメラマルチメディア新宿東口": {
			"ヨドバシカメラマルティメディア新宿東口",
		},
		"ヨドバシカメラマルチメディア町田": {
			"ヨドバシカメラマルチメディア町田駅前",
		},
		"レジャーランド高崎駅東口": {
			"レジャーランド高崎",
		},
		"わくわくカーニバル": {
			"わくわくカーニバル神戸",
		},
		"ワンダーフォレスト江南西": {
			"ハローズガーデンワンダーフォレスト江南西",
		},
	}
}

// AggregateShops returns aggregated shops with similar name
func AggregateShops(shops []*prismdb.Shop) []*prismdb.Shop {
	var reversedSimilarShopNames = map[string]string{}

	for key, values := range similarShopNames {
		foldedKey := width.Fold.String(key)
		for _, value := range values {
			foldedValue := width.Fold.String(value)
			reversedSimilarShopNames[foldedValue] = foldedKey
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

		foldedShopName := width.Fold.String(shopName)
		if reversedSimilarShopNames[foldedShopName] != "" {
			shopName = reversedSimilarShopNames[foldedShopName]
		}

		shopName = width.Fold.String(shopName)

		shopName = strings.ReplaceAll(shopName, "モーリーファンタジー・f", "モーリーファンタジーf")
		shopName = strings.ReplaceAll(shopName, "CLUBSEGA", "クラブセガ")
		shopName = strings.ReplaceAll(shopName, "ヤマダ電機LABI", "LABI")

		shopName = regexp.MustCompile(`([^A-Za-z0-9])\s+([^A-Za-z0-9])`).ReplaceAllString(shopName, "$1$2")
		shopName = regexp.MustCompile(`(?i)SOYU\s*Game\s*Field`).ReplaceAllString(shopName, "ソユーゲームフィールド")
		shopName = regexp.MustCompile(`^ニコパ`).ReplaceAllString(shopName, "NICOPA")

		shopName = strings.TrimSpace(shopName)

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