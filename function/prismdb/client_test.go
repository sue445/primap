package prismdb

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/sue445/primap/testutil"
	"testing"
)

func TestClient_GetAllShops(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://web-lk3h3ydj7a-an.a.run.app/sparql",
		httpmock.NewStringResponder(200, testutil.ReadTestData("../testdata/all_shops.json")))

	want := &Shop{
		Name:       "モーリーファンタジー伊勢ララパーク",
		Prefecture: "三重県",
		Address:    "三重県伊勢市小木町曽弥538 ｲｵﾝﾗﾗﾊﾟｰｸSC2階",
		Series:     []string{"primagi_1"},
	}

	c, err := NewClient()

	if !assert.NoError(t, err) {
		return
	}

	got, err := c.GetAllShops()

	if !assert.NoError(t, err) {
		return
	}

	if !assert.Greater(t, len(got), 0) {
		return
	}
	assert.Equal(t, want, got[0])
}
