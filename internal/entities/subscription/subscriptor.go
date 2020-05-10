package subscription

import (
	"github.com/ffernan01/minimq/internal/entities/topic"
	"github.com/google/uuid"
)

// SubscriptorID is the ID of the subscriptor
type SubscriptorID uuid.UUID

// Subscriptor is the specification for a subscriptor
type Subscriptor struct {
	ID          SubscriptorID
	Topic       topic.Topic
	Description string
}
