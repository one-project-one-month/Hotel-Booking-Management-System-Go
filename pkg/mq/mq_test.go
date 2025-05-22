package mq

import (
	"slices"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageQueue(t *testing.T) {
	var wg sync.WaitGroup
	mq := New(&wg, 10)
	topic := "test_topic"
	limit := 10

	subscriber := func(data any) any {
		return data
	}
	mq.Subscribe(topic, subscriber)

	wg.Add(1)
	result := make([]int, 0)
	go func() {
		defer wg.Done()
		for i := range limit {
			message := Message{
				AppID: "test_app",
				Topic: topic,
				Data:  i,
			}
			wg.Add(1)
			go func() {
				defer wg.Done()
				r := mq.Publish(&message)
				for data := range r {
					result = append(result, data.(int))
				}
			}()
		}
	}()

	wg.Wait()

	assert.Equal(t, limit, len(result), "Expected %d results, but got %d", limit, len(result))

	slices.Sort(result)
	for i := range limit {
		assert.Equal(t, i, result[i], "Expected %d, but got %d", i, result[i])
	}
}
