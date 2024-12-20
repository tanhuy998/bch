package memoryCache

import (
	"fmt"
	"sync"
	"time"
)

/*
memoryCache is a sync.Map that store "read/write lockable" rooms for cached value
in order to deal with thread safety concern between goroutine
*/
type (
	cache_vault struct {
		sync.Map
	}
)

var (
	cache_topics  *cache_vault
	cleanupTicker *time.Ticker
	endSignal     chan bool
)

var (
	ERR_TOPIC_EXIST              = fmt.Errorf("topic exists")
	ERR_INVALID_TOPIC_CONSTRAINT = fmt.Errorf("invalid topic constraint")
)

func init() {

	cache_topics = new(cache_vault)
	cleanupTicker = time.NewTicker(time.Minute * 2)

	//go poll()
}

func poll() {

	for {
		select {
		case <-cleanupTicker.C:
			cleanup()
		case <-endSignal:
			clearAll()
		}
	}
}

func cleanup() {

	cache_topics.Range(func(key interface{}, val interface{}) bool {

		if c, ok := val.(ISelfCleanupCacheUnit); ok {

			go c.cleanup()
			return true
		}

		return true
	})
}

func clearAll() {

	cache_topics = nil
}

func Terminate() {

	endSignal <- true
}

func GetTopic[Key_T, Value_T comparable](topic string) (*cache_unit[Key_T, Value_T], bool, error) {

	unknown, ok := cache_topics.Load(topic)

	if !ok {

		return nil, false, nil
	}

	if val, ok := unknown.(*cache_unit[Key_T, Value_T]); ok {

		return val, true, nil
	}

	return nil, true, ERR_INVALID_TOPIC_CONSTRAINT
}

func NewTopic[Key_T, Value_T comparable](topic string) error {

	if _, exists := cache_topics.Load(topic); exists {

		return ERR_TOPIC_EXIST
	}

	newCacheUnit := newCacheUnit[Key_T, Value_T](topic)
	newCacheUnit.topic = topic

	cache_topics.Store(topic, newCacheUnit)

	return nil
}

func DropTopic(topic string) {

	cache_topics.LoadAndDelete(topic)
}
