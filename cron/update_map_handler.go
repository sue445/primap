package cron

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/sue445/primap/db"
	"github.com/sue445/primap/prismdb"
	"log"
	"net/http"
)

// UpdateMapHandler returns handler of /cron/update_map
func UpdateMapHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Impl

	fmt.Fprint(w, "ok")
}

func handleError(w http.ResponseWriter, err error) {
	log.Printf("[ERROR] %+v", err)
	sentry.CaptureException(err)
	w.WriteHeader(500)
	fmt.Fprint(w, err)
}

func toEntity(shop *prismdb.Shop) *db.ShopEntity {
	return &db.ShopEntity{
		Name:       shop.Name,
		Prefecture: shop.Prefecture,
		Address:    shop.Address,
		Series:     shop.Series,
	}
}
