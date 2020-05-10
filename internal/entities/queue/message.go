package queue

import (
	"github.com/ffernan01/minimq/internal/entities/message"
	"github.com/ffernan01/minimq/internal/entities/topic"
	"github.com/google/uuid"
)

// MessageID is the queued message ID
type MessageID uuid.UUID

// Message is the enqueued message
type Message struct {
	ID      MessageID
	TopicID topic.ID
	Content message.Content
	QueueID ID
}
