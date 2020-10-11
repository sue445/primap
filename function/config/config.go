package config

import "os"

// InitParams defines args for Init()
type InitParams struct {
	ProjectID   string
	Environment string
}

var projectID string
var environment string

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

// GetEnvironment returns Environment in config
func GetEnvironment() string {
	return environment
}

// IsTest returns whether test env
func IsTest() bool {
	return GetProjectID() == "test"
}
