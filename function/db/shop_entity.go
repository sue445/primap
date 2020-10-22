package db

import (
	"context"
	"github.com/getsentry/sentry-go"
	"github.com/mmcloughlin/geohash"
	"github.com/pkg/errors"
	secretmanagerenv "github.com/sue445/gcp-secretmanagerenv"
	"github.com/sue445/primap/config"
	"golang.org/x/text/width"
	"google.golang.org/genproto/googleapis/type/latlng"
	"googlemaps.github.io/maps"
	"log"
	"regexp"
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

// UpdateAddressWithGeography update address and fetch geography if necessary
func (e *ShopEntity) UpdateAddressWithGeography(ctx context.Context, address string) error {
	sanitizedAddress := sanitizeAddress(address)

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

func sanitizeAddress(address string) string {
	sanitized := width.Fold.String(address)

	// Normalize Japanese street number(丁目,番地,号)
	sanitized = regexp.MustCompile(`([0-9]+)番地の([0-9]+)`).ReplaceAllString(sanitized, "$1-$2")
	sanitized = regexp.MustCompile(`([0-9]+)(?:(?:丁目)|(?:番地?)|(?:号))`).ReplaceAllString(sanitized, "$1-")

	// Remove building name after street name
	sanitized = regexp.MustCompile(`([0-9]+(?:-[0-9]+)?(?:-[0-9]+)?)[^条線]*$`).ReplaceAllString(sanitized, "$1")

	return sanitized
}
