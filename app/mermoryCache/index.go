package memoryCache

import (
	"errors"
	"sync"
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
	topics *cache_vault = new(cache_vault)
)

var (
	ERR_TOPIC_EXIST = errors.New("topic exists")
)

func GetTopic[Key_T, Value_T any](topic string) (*cache_unit[Key_T, *cache_value[Value_T]], bool) {

	unknown, ok := topics.Load(topic)

	if !ok {

		return nil, false
	}

	if val, ok := unknown.(*cache_unit[Key_T, *cache_value[Value_T]]); ok {

		return val, true
	}

	return nil, false
}

func NewTopic[Key_T, Value_T any](topic string) error {

	if _, exists := topics.Load(topic); exists {

		return ERR_TOPIC_EXIST
	}

	newCacheUnit := new(cache_unit[Key_T, cache_value[Value_T]])

	topics.Store(topic, newCacheUnit)

	return nil
}

func DeleteTopic(topic string) {

	topics.LoadAndDelete(topic)
}
