package mongoDBTracerService

import (
	accessLogServicePort "app/port/accessLog"
	"context"
	"fmt"
	"sync"
	"time"
)

type (
	MongoDBQueryDetail struct {
		QueryType  string  `json:"query_type"`
		Collection string  `json:"collection"`
		DBTime     float64 `json:"db_time_ms"`
		Err        error   `json:"error,omitempty"`
	}

	MongoDBTracerMonitor struct {
		Detail      []MongoDBQueryDetail `json:"detail"`
		TotalDBTime uint64               `json:"total_db_time_ms"`
	}
)

type (
	DBQueryTracerService struct {
		sync.Mutex
		AccessLogger accessLogServicePort.IAccessLogger[MongoDBTracerMonitor]
	}
)

func (this *DBQueryTracerService) Trace(collectionName string, label string, ctx context.Context) func(error) {
	fmt.Println("db trace")
	dbMonitor := this.AccessLogger.GetDBMonitor(ctx)

	begin := time.Now()

	return func(err error) {

		this.AccessLogger.AsyncTask(
			ctx,
			func() {

				end := time.Since(begin)
				this.Lock()
				defer this.Unlock()

				q := MongoDBQueryDetail{
					QueryType:  label,
					Collection: collectionName,
					DBTime:     float64(end) / float64(time.Millisecond),
					Err:        err,
				}

				dbMonitor.Detail = append(dbMonitor.Detail, q)
			},
		)
	}
}
