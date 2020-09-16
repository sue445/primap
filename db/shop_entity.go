package db

import (
	"google.golang.org/genproto/googleapis/type/latlng"
	"time"
)

const (
	shopCollectionName = "Shops"
)

// ShopEntity represents a shop entity for Firestore
type ShopEntity struct {
	Name       string         `firestore:"name"`
	Prefecture string         `firestore:"prefecture"`
	Address    string         `firestore:"address"`
	Series     []string       `firestore:"series"`
	CreatedAt  time.Time      `firestore:"created_at"`
	UpdatedAt  time.Time      `firestore:"updated_at"`
	Location   *latlng.LatLng `firestore:"latlng"`
}
