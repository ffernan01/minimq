package gateway

import (
	"github.com/ffernan01/minimq/internal/entities/subscription"
	"github.com/ffernan01/minimq/internal/entities/topic"
)

// SubscriptionsRepository repository for subscriptions
type SubscriptionsRepository interface {
	Add(*subscription.Subscriptor) error
	Remove(*subscription.Subscriptor) error
	GetByTopic(topic.ID) ([]subscription.Subscriptor, error)
}
