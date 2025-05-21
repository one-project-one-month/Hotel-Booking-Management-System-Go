package mq

import (
	"sync"
)

// Subscriber is a function type that processes message data and returns a result.
type Subscriber func(data any) any

// MQ represents a message queue that manages publishing and subscribing to topics.
type MQ struct {
	queue  chan *Message
	broker *broker
}

// Start begins processing messages from the queue in background.
func (mq *MQ) Start() {
	go func() {
		for message := range mq.queue {
			mq.broker.deliver(message)
		}
	}()
}

// Publish sends a message to the message queue.
func (mq *MQ) Publish(message *Message) chan any {
	reply := make(chan any)
	go func() {
		message.ReplyTo = reply
		mq.queue <- message
	}()
	return reply
}

// Subscribe registers a subscriber function to process messages for a specific topic.
func (mq *MQ) Subscribe(topic string, subscriber Subscriber) {
	mq.broker.receive(topic, subscriber)
}

// New creates and initializes a new message queue with the specified wait group and queue size.
func New(wg *sync.WaitGroup, size uint) *MQ {
	broker := broker{
		wg:          wg,
		mu:          sync.RWMutex{},
		subscribers: make(map[string][]Subscriber),
	}

	mq := &MQ{
		queue:  make(chan *Message, size),
		broker: &broker,
	}

	mq.Start()

	return mq
}
