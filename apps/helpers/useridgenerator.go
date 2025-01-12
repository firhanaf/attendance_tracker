package helpers

import "github.com/google/uuid"

func UserIDGenerator() string {
	return uuid.NewString()
}
