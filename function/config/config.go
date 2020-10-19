package config

import "os"

// InitParams defines args for Init()
type InitParams struct {
	ProjectID string
}

var projectID string

// Init Setups config
func Init(args *InitParams) {
	projectID = args.ProjectID
}

// GetProjectID returns ProjectID in config
func GetProjectID() string {
	if projectID != "" {
		return projectID
	}

	return os.Getenv("GCP_PROJECT")
}

// IsTest returns whether test env
func IsTest() bool {
	return GetProjectID() == "test"
}
