package db

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sue445/primap/server/config"
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
	Name       string         `firestore:"name"       json:"name"`
	Prefecture string         `firestore:"prefecture" json:"prefecture"`
	Address    string         `firestore:"address"    json:"address"`
	Series     []string       `firestore:"series"     json:"series"`
	CreatedAt  time.Time      `firestore:"created_at" json:"created_at"`
	UpdatedAt  time.Time      `firestore:"updated_at" json:"updated_at"`
	Location   *latlng.LatLng `firestore:"location"   json:"location"`
	Latitude   float64        `firestore:"latitude"   json:"latitude"`
	Longitude  float64        `firestore:"longitude"  json:"longitude"`
	Deleted    bool           `firestore:"deleted"    json:"deleted"`
}

// UpdateAddressWithLocation update address and fetch location if necessary
func (e *ShopEntity) UpdateAddressWithLocation(ctx context.Context, address string) error {
	if e.Address == address && e.Location != nil {
		return nil
	}

	// If Address is changed, should update Location.
	// But when Location is nil(undefined), always should update.

	if config.GetGoogleMapsAPIKey() != "" {
		c, err := maps.NewClient(maps.WithAPIKey(config.GetGoogleMapsAPIKey()))

		if err != nil {
			return errors.WithStack(err)
		}

		r := &maps.GeocodingRequest{Address: address}
		resp, err := c.Geocode(ctx, r)

		if err != nil {
			return errors.WithStack(err)
		}

		if len(resp) > 0 {
			e.Location = &latlng.LatLng{
				Latitude:  resp[0].Geometry.Location.Lat,
				Longitude: resp[0].Geometry.Location.Lng,
			}
			e.Latitude = resp[0].Geometry.Location.Lat
			e.Longitude = resp[0].Geometry.Location.Lng
		} else {
			log.Printf("[WARN] Location is unknown: Address=%s, Shop=%+v", address, e)
			e.Location = nil
		}
	}

	e.Address = address
	return nil
}
