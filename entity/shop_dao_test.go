package entity

import (
	"github.com/stretchr/testify/assert"
	"github.com/sue445/primap/testutil"
	"testing"
)

func Test_sliceShops_Sliced(t *testing.T) {
	shops := []*ShopEntity{
		{
			Name: "shop1",
		},
		{
			Name: "shop2",
		},
		{
			Name: "shop3",
		},
		{
			Name: "shop4",
		},
		{
			Name: "shop5",
		},
	}

	got := sliceShops(shops, 2)

	assert.Equal(t, &ShopEntity{Name: "shop1"}, got[0][0])
	assert.Equal(t, &ShopEntity{Name: "shop2"}, got[0][1])
	assert.Equal(t, &ShopEntity{Name: "shop3"}, got[1][0])
	assert.Equal(t, &ShopEntity{Name: "shop4"}, got[1][1])
	assert.Equal(t, &ShopEntity{Name: "shop5"}, got[2][0])
}

func Test_sliceShops_NotSliced1(t *testing.T) {
	shops := []*ShopEntity{
		{
			Name: "shop1",
		},
		{
			Name: "shop2",
		},
	}

	got := sliceShops(shops, 2)

	assert.Equal(t, [][]*ShopEntity{shops}, got)
}

func Test_sliceShops_NotSliced2(t *testing.T) {
	shops := []*ShopEntity{
		{
			Name: "shop1",
		},
	}

	got := sliceShops(shops, 2)

	assert.Equal(t, [][]*ShopEntity{shops}, got)
}

func TestShopDao_SaveShops_And_GetShop(t *testing.T) {
	defer testutil.CleanupFirestore()

	revision := "20200123-123456"

	shop := &ShopEntity{
		Name:       "ＭＥＧＡドン・キホーテＵＮＹ名張",
		Prefecture: "三重県",
		Address:    "三重県名張市下比奈知黒田3100番地の1",
		Series:     []string{"prichan"},
	}

	dao := NewShopDao(testutil.TestProjectID)
	err := dao.SaveShops([]*ShopEntity{shop}, revision)

	if !assert.NoError(t, err) {
		return
	}

	got1, err := dao.GetShop("ＭＥＧＡドン・キホーテＵＮＹ名張")

	if assert.NoError(t, err) {
		assert.Equal(t, "ＭＥＧＡドン・キホーテＵＮＹ名張", got1.Name)
		assert.Equal(t, "三重県", got1.Prefecture)
		assert.Equal(t, "三重県名張市下比奈知黒田3100番地の1", got1.Address)
		assert.Equal(t, "20200123-123456", got1.Revision)
		assert.Equal(t, []string{"prichan"}, got1.Series)
		assert.NotNil(t, got1.CreatedAt)
		assert.NotNil(t, got1.UpdatedAt)
	}

	got2, err := dao.GetShop("UNKNOWN")

	if assert.NoError(t, err) {
		assert.Nil(t, got2)
	}
}
