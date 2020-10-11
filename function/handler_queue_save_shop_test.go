package primap

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/sue445/primap/config"
	"github.com/sue445/primap/db"
	"github.com/sue445/primap/prismdb"
	"github.com/sue445/primap/testutil"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	config.Init(&config.InitParams{
		Environment: "test",
	})

	code := m.Run()

	config.Init(&config.InitParams{
		Environment: "",
	})

	os.Exit(code)
}

func Test_saveShop(t *testing.T) {
	shop := &prismdb.Shop{
		Name:       "ＭＥＧＡドン・キホーテＵＮＹ名張",
		Prefecture: "三重県",
		Address:    "三重県名張市下比奈知黒田3100番地の1",
		Series:     []string{"prichan"},
	}

	ctx := context.Background()
	projectID := testutil.TestProjectID()
	err := saveShop(ctx, projectID, shop)

	if !assert.NoError(t, err) {
		return
	}

	dao := db.NewShopDao(ctx, projectID)
	got, err := dao.LoadShop("ＭＥＧＡドン・キホーテＵＮＹ名張")

	if assert.NoError(t, err) {
		if assert.NotNil(t, got) {
			assert.Equal(t, "ＭＥＧＡドン・キホーテＵＮＹ名張", got.Name)
			assert.Equal(t, "三重県", got.Prefecture)
			assert.Equal(t, "三重県名張市下比奈知黒田3100番地の1", got.Address)
			assert.Equal(t, []string{"prichan"}, got.Series)
			assert.False(t, got.CreatedAt.IsZero())
			assert.False(t, got.UpdatedAt.IsZero())
			assert.Nil(t, got.Geography)
		}
	}
}

func Test_saveShop_DeletedShipWillReborn(t *testing.T) {
	shop := &prismdb.Shop{
		Name:       "ＭＥＧＡドン・キホーテＵＮＹ名張",
		Prefecture: "三重県",
		Address:    "三重県名張市下比奈知黒田3100番地の1",
		Series:     []string{"prichan"},
	}

	projectID := testutil.TestProjectID()
	ctx := context.Background()

	dao := db.NewShopDao(ctx, projectID)
	err := dao.SaveShop(
		&db.ShopEntity{
			Name:      "ＭＥＧＡドン・キホーテＵＮＹ名張",
			CreatedAt: time.Now(),
			Deleted:   true,
		},
	)
	if !assert.NoError(t, err) {
		return
	}

	err = saveShop(ctx, projectID, shop)

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
			assert.False(t, got.CreatedAt.IsZero())
			assert.False(t, got.UpdatedAt.IsZero())
			assert.Nil(t, got.Geography)
			assert.False(t, got.Deleted)
		}
	}
}
