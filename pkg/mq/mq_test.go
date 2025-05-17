package mq

import (
	"slices"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMessageQueue(t *testing.T) {
	var wg sync.WaitGroup
	mq := New(&wg, 10)
	topic := "test_topic"
	reply := make(chan any)
	limit := 100

	subscriber := func(data any) any {
		return data
	}
	mq.Subscribe(topic, subscriber)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range limit {
			message := Message{
				AppID:   "test_app",
				Topic:   topic,
				Data:    i,
				ReplyTo: reply,
			}
			mq.Publish(&message)
		}
	}()

	result := make([]int, limit)
loop:
	for i := range limit {
		select {
		case data := <-reply:
			result[i] = data.(int)
		case <-time.Tick(time.Second):
			t.Error("Expected a reply, but got none")
			break loop
		}
	}

	wg.Wait()

	assert.Equal(t, limit, len(result), "Expected %d results, but got %d", limit, len(result))

	slices.Sort(result)
	for i := range limit {
		assert.Equal(t, i, result[i], "Expected %d, but got %d", i, result[i])
	}

	close(reply)
}
