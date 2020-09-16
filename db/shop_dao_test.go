package db

import (
	"github.com/stretchr/testify/assert"
	"github.com/sue445/primap/testutil"
	"testing"
)

func TestShopDao_SaveShop_And_LoadShop(t *testing.T) {
	defer testutil.CleanupFirestore()

	shop := &ShopEntity{
		Name:       "ＭＥＧＡドン・キホーテＵＮＹ名張",
		Prefecture: "三重県",
		Address:    "三重県名張市下比奈知黒田3100番地の1",
		Series:     []string{"prichan"},
	}

	dao := NewShopDao(testutil.TestProjectID())
	err := dao.SaveShop(shop)

	if !assert.NoError(t, err) {
		return
	}

	got1, err := dao.LoadShop("ＭＥＧＡドン・キホーテＵＮＹ名張")

	if assert.NoError(t, err) {
		if assert.NotNil(t, got1) {
			assert.Equal(t, "ＭＥＧＡドン・キホーテＵＮＹ名張", got1.Name)
			assert.Equal(t, "三重県", got1.Prefecture)
			assert.Equal(t, "三重県名張市下比奈知黒田3100番地の1", got1.Address)
			assert.Equal(t, []string{"prichan"}, got1.Series)
			assert.NotNil(t, got1.UpdatedAt)
		}
	}

	got2, err := dao.LoadShop("UNKNOWN")

	if assert.NoError(t, err) {
		assert.Nil(t, got2)
	}
}
