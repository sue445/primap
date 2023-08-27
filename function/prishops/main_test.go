package prishops

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/sue445/primap/testutil"
	"testing"
)

func TestGetAllShops(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://web-lk3h3ydj7a-an.a.run.app/sparql",
		httpmock.NewStringResponder(200, testutil.ReadTestData("../testdata/all_shops.json")))

	gotShops, err := GetAllShops()

	if !assert.NoError(t, err) {
		return
	}

	if !assert.Greater(t, len(gotShops), 0) {
		return
	}

	wantShop := &Shop{
		Name:       "モーリーファンタジー伊勢ララパーク",
		Prefecture: "三重県",
		Address:    "三重県伊勢市小木町曽弥538 ｲｵﾝﾗﾗﾊﾟｰｸSC2階",
		Series:     []string{"primagi_1"},
	}
	assert.Equal(t, wantShop, gotShops[0])

	wantShop2 := &Shop{
		Name:       "プリズムストーンカフェ原宿店",
		Prefecture: "東京都",
		Address:    "東京都渋谷区神宮前3-18-27 2F",
		Series:     []string{"primagi_1", "prismstone", "pripara"},
	}
	assert.Equal(t, wantShop2, gotShops[1016])
}
