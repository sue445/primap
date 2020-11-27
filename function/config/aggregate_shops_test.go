package config

import (
	"github.com/stretchr/testify/assert"
	"github.com/sue445/primap/prismdb"
	"testing"
)

func TestAggregateShops(t *testing.T) {
	shops := []*prismdb.Shop{
		{
			Name:       "ふぇすたらんど小野店",
			Prefecture: "兵庫県",
			Address:    "兵庫県小野市王子町８６８－１ イオン小野店２Ｆ",
			Series:     []string{"pripara"},
		},
		{
			Name:       "ふぇすたらんど小野",
			Prefecture: "兵庫県",
			Address:    "兵庫県小野市王子町８６８－１イオン小野店２Ｆ",
			Series:     []string{"prichan"},
		},
		{
			Name:       "モーリーファンタジー唐津",
			Prefecture: "佐賀県",
			Address:    "佐賀県唐津市鏡字立神４６７１ イオン２階",
			Series:     []string{"prichan", "pripara"},
		},
		{
			Name:       "プリズムストーン  新札幌",
			Prefecture: "北海道",
			Address:    "北海道札幌市厚別区厚別中央２条５丁目６－２　ＤＵＯ－１　４階－１イオン小野店２Ｆ",
			Series:     []string{"prichan"},
		},
		{
			Name:       "プリズムストーン札幌",
			Prefecture: "北海道",
			Address:    "北海道札幌市厚別区厚別中央２条５丁目６－２ ＤＵＯ－１ ４階",
			Series:     []string{"pripara"},
		},
		{
			Name:       "ＳＯＹＵＧａｍｅＦｉｅｌｄ湘南",
			Prefecture: "神奈川県",
			Address:    "神奈川県藤沢市辻堂新町四丁目１番１号 湘南モールＦＩＬＬ２Ｆ",
			Series:     []string{"pripara"},
		},
		{
			Name:       "ＳＯＹＵ　Ｇａｍｅ　Ｆｉｅｌｄ湘南店",
			Prefecture: "神奈川県",
			Address:    "神奈川県藤沢市辻堂新町四丁目１番１号　湘南モールＦＩＬＬ２Ｆ",
			Series:     []string{"prichan"},
		},
	}

	got := AggregateShops(shops)

	want := []*prismdb.Shop{
		{
			Name:       "ふぇすたらんど小野",
			Prefecture: "兵庫県",
			Address:    "兵庫県小野市王子町８６８－１ イオン小野店２Ｆ",
			Series:     []string{"prichan", "pripara"},
		},
		{
			Name:       "プリズムストーン札幌",
			Prefecture: "北海道",
			Address:    "北海道札幌市厚別区厚別中央２条５丁目６－２　ＤＵＯ－１　４階－１イオン小野店２Ｆ",
			Series:     []string{"prichan", "pripara"},
		},
		{
			Name:       "モーリーファンタジー唐津",
			Prefecture: "佐賀県",
			Address:    "佐賀県唐津市鏡字立神４６７１ イオン２階",
			Series:     []string{"prichan", "pripara"},
		},
		{
			Name:       "ＳＯＹＵＧａｍｅＦｉｅｌｄ湘南",
			Prefecture: "神奈川県",
			Address:    "神奈川県藤沢市辻堂新町四丁目１番１号 湘南モールＦＩＬＬ２Ｆ",
			Series:     []string{"prichan", "pripara"},
		},
	}

	assert.Equal(t, want, got)
}
