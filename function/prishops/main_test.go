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
		Name:       "ＭＥＧＡドン・キホーテＵＮＹ名張",
		Prefecture: "三重県",
		Address:    "三重県名張市下比奈知黒田３１００番地の１",
		Series:     []string{"prichan"},
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