package db

import "time"

const (
	shopCollectionName = "Shops"
)

// ShopEntity represents a shop entity for Firestore
type ShopEntity struct {
	Name       string    `firestore:"name"`
	Prefecture string    `firestore:"prefecture"`
	Address    string    `firestore:"address"`
	Revision   string    `firestore:"revision"`
	Series     []string  `firestore:"series"`
	UpdatedAt  time.Time `firestore:"updated_at"`
}
