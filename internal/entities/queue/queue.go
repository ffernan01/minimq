package queue

import (
	"github.com/ffernan01/minimq/internal/entities/message"
	"github.com/ffernan01/minimq/internal/entities/topic"
	"github.com/google/uuid"
)

// ID is the Id for the queue
type ID uuid.UUID

// Queue has the enqueued messages for a specific topic
type Queue struct {
	ID    ID
	Topic topic.Topic
}

// Enqueue a message
func (q *Queue) Enqueue(topic topic.Topic, message *message.Message) (*Message, error) {
	return nil, nil
}

// Peek takes the first message in the queue without dequeuing it
func (q *Queue) Peek() (*Message, error) {
	return nil, nil
}

// Dequeue the first message in the queue
func (q *Queue) Dequeue() (*Message, error) {
	return nil, nil
}
