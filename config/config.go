package config

// InitParams defines args for Init()
type InitParams struct {
	GoogleMapsAPIKey string
}

var googleMapsAPIKey string

// Init Setups config
func Init(args *InitParams) {
	googleMapsAPIKey = args.GoogleMapsAPIKey
}

// GetGoogleMapsAPIKey returns GoogleMapsAPIKey in config
func GetGoogleMapsAPIKey() string {
	return googleMapsAPIKey
}
