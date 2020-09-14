package cron

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/sue445/primap/db"
	"github.com/sue445/primap/testutil"
	"testing"
	"time"
)

func Test_updateMap(t *testing.T) {
	// defer testutil.CleanupFirestore()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://prismdb.takanakahiko.me/sparql",
		httpmock.NewStringResponder(200, testutil.ReadTestData("../testdata/all_shops.json")))

	time := time.Date(2020, 1, 23, 12, 34, 56, 0, time.UTC)

	err := updateMap(time)

	if !assert.NoError(t, err) {
		return
	}

	dao := db.NewShopDao(testutil.TestProjectID())
	got, err := dao.GetShop("ＭＥＧＡドン・キホーテＵＮＹ名張")

	if assert.NoError(t, err) {
		if assert.NotNil(t, got) {
			assert.Equal(t, "ＭＥＧＡドン・キホーテＵＮＹ名張", got.Name)
			assert.Equal(t, "三重県", got.Prefecture)
			assert.Equal(t, "三重県名張市下比奈知黒田3100番地の1", got.Address)
			assert.Equal(t, "20200123-123456", got.Revision)
			assert.Equal(t, []string{"prichan"}, got.Series)
			assert.NotNil(t, got.CreatedAt)
			assert.NotNil(t, got.UpdatedAt)
		}
	}
}
