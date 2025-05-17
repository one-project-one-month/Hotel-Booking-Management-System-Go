package mq

// Message represents a message that can be published to the message queue.
type Message struct {
	AppID   string
	Topic   string
	Data    any
	ReplyTo chan any
}
