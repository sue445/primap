package db

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sue445/primap/config"
	"google.golang.org/genproto/googleapis/type/latlng"
	"googlemaps.github.io/maps"
	"log"
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
			return errors.WithStack(err)
		}

		r := &maps.GeocodingRequest{Address: address}
		resp, err := c.Geocode(context.Background(), r)

		if err != nil {
			return errors.WithStack(err)
		}

		if len(resp) > 0 {
			e.Location = &latlng.LatLng{
				Latitude:  resp[0].Geometry.Location.Lat,
				Longitude: resp[0].Geometry.Location.Lng,
			}
		} else {
			log.Printf("[WARN] Location is unknown: Address=%s, Shop=%+v", address, e)
			e.Location = nil
		}
	}

	e.Address = address
	return nil
}
