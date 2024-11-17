package memoryCache

type (
	CacheTopicSnapShot[Key_T, Value_T comparable] struct {
		Topic string `json:"topic"`
	}
)
