package commons

import "github.com/google/uuid"

func GenerateUUID() string {
	id := uuid.New()
	return id.String()
}
