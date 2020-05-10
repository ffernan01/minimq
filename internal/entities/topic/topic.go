package topic

import "github.com/google/uuid"

// ID is the ID type for topics
type ID uuid.UUID

// Topic of a message
type Topic struct {
	ID   ID
	Name string
}
