package message

// Content is the content of a message
type Content string

// Message is the base unit of the MQ System
type Message struct {
	Content Content
}
