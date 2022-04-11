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

	want := &Shop{
		Name:       "NICOPAイオン伊勢",
		Prefecture: "三重県",
		Address:    "三重県伊勢市楠部町乙160 ｲｵﾝ伊勢店",
		Series:     []string{"primagi"},
	}

	got, err := GetAllShops()

	if !assert.NoError(t, err) {
		return
	}

	if !assert.Greater(t, len(got), 0) {
		return
	}
	assert.Equal(t, want, got[0])

	assert.Contains(t, got, shopList[0])
}
