package cron

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/itchyny/timefmt-go"
	"github.com/sue445/primap/db"
	"github.com/sue445/primap/prismdb"
	"log"
	"net/http"
	"os"
	"time"
)

// UpdateMapHandler returns handler of /cron/update_map
func UpdateMapHandler(w http.ResponseWriter, r *http.Request) {
	err := updateMap(time.Now())

	if err != nil {
		handleError(w, err)
		return
	}

	fmt.Fprint(w, "ok")
}

func handleError(w http.ResponseWriter, err error) {
	log.Printf("[ERROR] %+v", err)
	sentry.CaptureException(err)
	w.WriteHeader(500)
	fmt.Fprint(w, err)
}

func updateMap(time time.Time) error {
	client, err := prismdb.NewClient()
	if err != nil {
		return err
	}

	shops, err := client.GetAllShops()
	if err != nil {
		return err
	}

	var entityShops []*db.ShopEntity
	for _, shop := range shops {
		entityShops = append(entityShops, toEntity(shop))
	}

	revision := timefmt.Format(time, "%Y%m%d-%H%M%S")

	projectID := os.Getenv("GCP_PROJECT")
	dao := db.NewShopDao(projectID)

	err = dao.SaveShops(entityShops, revision)
	if err != nil {
		return err
	}

	log.Printf("Successful: revision=%s, shops=%d\n", revision, len(entityShops))

	return nil
}

func toEntity(shop *prismdb.Shop) *db.ShopEntity {
	return &db.ShopEntity{
		Name:       shop.Name,
		Prefecture: shop.Prefecture,
		Address:    shop.Address,
		Series:     shop.Series,
	}
}
