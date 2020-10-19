package config

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
	return projectID
}
