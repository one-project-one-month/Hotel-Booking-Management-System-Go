// Package mq provides a message queue implementation for asynchronous communication between components.
package mq

import (
	"sync"
)

type broker struct {
	mu          sync.RWMutex
	wg          *sync.WaitGroup
	subscribers map[string][]Subscriber
}

func (b *broker) receive(topic string, s Subscriber) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.subscribers[topic] = append(b.subscribers[topic], s)
}

func (b *broker) deliver(msg *Message) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if subscribers, found := b.subscribers[msg.Topic]; found {
		for _, subscriber := range subscribers {
			b.wg.Add(1)
			go func() {
				defer b.wg.Done()
				resp := subscriber(msg.Data)
				msg.ReplyTo <- resp
			}()
		}
	}
}
