package uuid

import "github.com/google/uuid"

// GenerateGUID is for generating Google Unique ID
func GenerateGUID() string {
	return uuid.NewString()
}
