package config

import "os"

// InitParams defines args for Init()
type InitParams struct {
	ProjectID        string
	GoogleMapsAPIKey string
}

var projectID string
var googleMapsAPIKey string

// Init Setups config
func Init(args *InitParams) {
	projectID = args.ProjectID
	googleMapsAPIKey = args.GoogleMapsAPIKey
}

// GetProjectID returns ProjectID in config
func GetProjectID() string {
	if projectID != "" {
		return projectID
	}

	return os.Getenv("GCP_PROJECT")
}

// GetGoogleMapsAPIKey returns GoogleMapsAPIKey in config
func GetGoogleMapsAPIKey() string {
	return googleMapsAPIKey
}
