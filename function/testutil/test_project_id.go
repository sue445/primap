package testutil

import "github.com/google/uuid"

// TestProjectID generate and return projectID for test
func TestProjectID() string {
	return uuid.New().String()
}
