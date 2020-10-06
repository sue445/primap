package db

import (
	"context"
	"github.com/mmcloughlin/geohash"
	"github.com/pkg/errors"
	"github.com/sue445/primap/config"
	"google.golang.org/genproto/googleapis/type/latlng"
	"googlemaps.github.io/maps"
	"log"
	"time"
)

const (
	shopCollectionName = "Shops"

	// c.f. https://github.com/codediodeio/geofirex#pointlatitude-number-longitude-number-firepoint
	geohashPrecision = 9
)

// ShopEntity represents a shop entity for Firestore
type ShopEntity struct {
	Name       string     `firestore:"name"       json:"name"`
	Prefecture string     `firestore:"prefecture" json:"prefecture"`
	Address    string     `firestore:"address"    json:"address"`
	Series     []string   `firestore:"series"     json:"series"`
	CreatedAt  time.Time  `firestore:"created_at" json:"created_at"`
	UpdatedAt  time.Time  `firestore:"updated_at" json:"updated_at"`
	Geography  *Geography `firestore:"geography"  json:"geography"`
	Deleted    bool       `firestore:"deleted"    json:"deleted"`
}

// UpdateAddressWithGeography update address and fetch geography if necessary
func (e *ShopEntity) UpdateAddressWithGeography(ctx context.Context, address string) error {
	if e.Address == address && e.Geography != nil {
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
			lat := resp[0].Geometry.Location.Lat
			lng := resp[0].Geometry.Location.Lng
			e.Geography = &Geography{
				GeoPoint: &latlng.LatLng{
					Latitude:  lat,
					Longitude: lng,
				},
				GeoHash: geohash.EncodeWithPrecision(lat, lng, geohashPrecision),
			}
		} else {
			log.Printf("[WARN] Location is unknown: Address=%s, Shop=%+v", address, e)
			e.Geography = nil
		}
	}

	e.Address = address
	return nil
}
