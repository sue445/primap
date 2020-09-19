package db

import (
	"context"
	"github.com/sue445/primap/config"
	"google.golang.org/genproto/googleapis/type/latlng"
	"googlemaps.github.io/maps"
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
	Deleted    bool           `firestore:"deleted"`
}

// UpdateAddressWithLocation update address and fetch location if necessary
func (e *ShopEntity) UpdateAddressWithLocation(address string) error {
	if e.Address == address {
		return nil
	}

	if config.GetGoogleMapsAPIKey() != "" {
		c, err := maps.NewClient(maps.WithAPIKey(config.GetGoogleMapsAPIKey()))

		if err != nil {
			return err
		}

		r := &maps.GeocodingRequest{Address: address}
		resp, err := c.Geocode(context.Background(), r)

		if err != nil {
			return err
		}

		e.Location = &latlng.LatLng{
			Latitude:  resp[0].Geometry.Location.Lat,
			Longitude: resp[0].Geometry.Location.Lng,
		}
	}

	e.Address = address
	return nil
}
