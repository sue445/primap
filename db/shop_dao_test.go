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

	got, err := dao.LoadShop("ＭＥＧＡドン・キホーテＵＮＹ名張")

	if assert.NoError(t, err) {
		if assert.NotNil(t, got) {
			assert.Equal(t, "ＭＥＧＡドン・キホーテＵＮＹ名張", got.Name)
			assert.Equal(t, "三重県", got.Prefecture)
			assert.Equal(t, "三重県名張市下比奈知黒田3100番地の1", got.Address)
			assert.Equal(t, []string{"prichan"}, got.Series)
			assert.True(t, got.CreatedAt.IsZero())
			assert.False(t, got.UpdatedAt.IsZero())
		}
	}
}

func TestShopDao_LoadShop(t *testing.T) {
	dao := NewShopDao(testutil.TestProjectID())

	got, err := dao.LoadShop("UNKNOWN")

	if assert.NoError(t, err) {
		assert.Nil(t, got)
	}
}

func TestShopDao_LoadOrCreateShop(t *testing.T) {
	dao := NewShopDao(testutil.TestProjectID())

	got, err := dao.LoadOrCreateShop("UNKNOWN")

	if assert.NoError(t, err) {
		if assert.NotNil(t, got) {
			assert.Equal(t, "UNKNOWN", got.Name)
			assert.Equal(t, "", got.Prefecture)
			assert.Equal(t, "", got.Address)
			assert.Equal(t, []string{}, got.Series)
			assert.False(t, got.CreatedAt.IsZero())
			assert.True(t, got.UpdatedAt.IsZero())
		}
	}
}
