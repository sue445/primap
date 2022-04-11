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

	httpmock.RegisterResponder("POST", "https://prismdb.takanakahiko.me/sparql",
		httpmock.NewStringResponder(200, testutil.ReadTestData("../testdata/all_shops.json")))

	gotShops, err := GetAllShops()

	if !assert.NoError(t, err) {
		return
	}

	if !assert.Greater(t, len(gotShops), 0) {
		return
	}

	wantShop := &Shop{
		Name:       "NICOPAイオン伊勢",
		Prefecture: "三重県",
		Address:    "三重県伊勢市楠部町乙160 ｲｵﾝ伊勢店",
		Series:     []string{"primagi"},
	}
	assert.Equal(t, wantShop, gotShops[0])

	wantShop2 := &Shop{
		Name:       "プリズムストーンカフェ原宿店",
		Prefecture: "東京都",
		Address:    "東京都渋谷区神宮前3-18-27 2F",
		Series:     []string{"primagi", "prismstone", "pripara"},
	}
	assert.Equal(t, wantShop2, gotShops[983])
}
