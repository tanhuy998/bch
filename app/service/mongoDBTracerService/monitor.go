package mongoDBTracerService

import (
	"time"
)

type (
	MongoDBQueryDetail struct {
		Label      string        `json:"operation"`
		DBType     string        `json:"db_type"`
		QueryType  string        `json:"query_type"`
		Collection string        `json:"collection"`
		DBTime     float64       `json:"db_time_ms"`
		dbTimeDur  time.Duration `json:"-"`
		Err        error         `json:"error,omitempty"`
	}
)

func (this *MongoDBQueryDetail) GetDBDuration() time.Duration {

	return this.dbTimeDur
}
