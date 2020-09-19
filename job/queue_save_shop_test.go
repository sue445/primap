package job

import (
	"github.com/stretchr/testify/assert"
	"github.com/sue445/primap/db"
	"github.com/sue445/primap/prismdb"
	"github.com/sue445/primap/testutil"
	"testing"
)

func Test_saveShop(t *testing.T) {
	testutil.SetRandomProjectID()
	// defer testutil.CleanupFirestore()

	shop := &prismdb.Shop{
		Name:       "ＭＥＧＡドン・キホーテＵＮＹ名張",
		Prefecture: "三重県",
		Address:    "三重県名張市下比奈知黒田3100番地の1",
		Series:     []string{"prichan"},
	}

	err := saveShop(testutil.TestProjectID(), shop)

	if !assert.NoError(t, err) {
		return
	}

	dao := db.NewShopDao(testutil.TestProjectID())
	got, err := dao.LoadShop("ＭＥＧＡドン・キホーテＵＮＹ名張")

	if assert.NoError(t, err) {
		if assert.NotNil(t, got) {
			assert.Equal(t, "ＭＥＧＡドン・キホーテＵＮＹ名張", got.Name)
			assert.Equal(t, "三重県", got.Prefecture)
			assert.Equal(t, "三重県名張市下比奈知黒田3100番地の1", got.Address)
			assert.Equal(t, []string{"prichan"}, got.Series)
			assert.False(t, got.CreatedAt.IsZero())
			assert.False(t, got.UpdatedAt.IsZero())
			assert.Nil(t, got.Location)
		}
	}
}
