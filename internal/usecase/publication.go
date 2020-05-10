package usecase

/*
import (
	"errors"

	"github.com/ffernan01/minimq/internal/entities/message"
	"github.com/ffernan01/minimq/internal/entities/queue"
	"github.com/ffernan01/minimq/internal/entities/topic"
)

// Publisher interface for publishing messages
type Publisher interface {
	Publish(topic.Topic, *message.Message) error
}

// PublisherImpl is a basic implementation of Publisher interface
type PublisherImpl struct {
	queueManager *queue.Manager
}

func NewPublisherImpl(queueManager *queue.Manager) (*PublisherImpl, error) {
	if queueManager == nil {
		return nil, errors.New("Undefined queueManager")
	}

	publisher := PublisherImpl{queueManager}
	return &publisher, nil
}

func (p *PublisherImpl) Publish(topic topic.Topic, message *message.Message) error {
	_, _, err := (*p.queueManager).Enqueue(topic, message)
	return err
}
*/
