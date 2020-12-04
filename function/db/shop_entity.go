package db

import (
	"context"
	"github.com/getsentry/sentry-go"
	"github.com/mmcloughlin/geohash"
	"github.com/pkg/errors"
	secretmanagerenv "github.com/sue445/gcp-secretmanagerenv"
	"github.com/sue445/primap/config"
	"github.com/sue445/primap/prismdb"
	"github.com/sue445/primap/util"
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
	Name             string     `firestore:"name"              json:"name"`
	Prefecture       string     `firestore:"prefecture"        json:"prefecture"`
	Address          string     `firestore:"address"           json:"address"`
	SanitizedAddress string     `firestore:"sanitized_address" json:"sanitized_address"`
	Series           []string   `firestore:"series"            json:"series"`
	CreatedAt        time.Time  `firestore:"created_at"        json:"created_at"`
	UpdatedAt        time.Time  `firestore:"updated_at"        json:"updated_at"`
	Geography        *Geography `firestore:"geography"         json:"geography"`
	Deleted          bool       `firestore:"deleted"           json:"deleted"`
}

// IsUpdated returns whether it has changed compared to prismdb.Shop
func (e *ShopEntity) IsUpdated(target *prismdb.Shop) bool {
	if e.Deleted || e.Geography == nil ||
		e.Name != target.Name || e.Prefecture != target.Prefecture ||
		e.Address != target.Address || e.SanitizedAddress != util.SanitizeAddress(target.Address) {

		return true
	}

	if len(e.Series) != len(target.Series) {
		return true
	}

	sourceSeries := util.SortedSlice(e.Series)
	targetSeries := util.SortedSlice(target.Series)

	for i := 0; i < len(sourceSeries); i++ {
		if sourceSeries[i] != targetSeries[i] {
			return true
		}
	}

	return false
}

// UpdateAddressWithGeography update address and fetch geography if necessary
func (e *ShopEntity) UpdateAddressWithGeography(ctx context.Context, address string) error {
	sanitizedAddress := util.SanitizeAddress(address)

	if e.Address == address && e.SanitizedAddress == sanitizedAddress && e.Geography != nil {
		return nil
	}

	// If Address is changed, should update Location.
	// But when Location is nil(undefined), always should update.

	googleMapsAPIKey, err := getGoogleMapsAPIKey(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	if googleMapsAPIKey != "" {
		c, err := maps.NewClient(maps.WithAPIKey(googleMapsAPIKey))

		if err != nil {
			return errors.WithStack(err)
		}

		r := &maps.GeocodingRequest{Address: sanitizedAddress}
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
			log.Printf("[WARN] Location is unknown: sanitizedAddress=%s, Shop=%+v", sanitizedAddress, e)

			sentry.ConfigureScope(func(scope *sentry.Scope) {
				scope.SetLevel(sentry.LevelWarning)
				scope.SetTag("sanitizedAddress", sanitizedAddress)
			})
			sentry.CaptureMessage("Location is unknown")

			e.Geography = nil
		}
	}

	e.Address = address
	e.SanitizedAddress = sanitizedAddress
	return nil
}

func getGoogleMapsAPIKey(ctx context.Context) (string, error) {
	secretmanager, err := secretmanagerenv.NewClient(ctx, config.GetProjectID())
	if err != nil {
		return "", errors.WithStack(err)
	}

	googleMapsAPIKey, err := secretmanager.GetValueFromEnvOrSecretManager("GOOGLE_MAPS_API_KEY", false)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return googleMapsAPIKey, nil
}
