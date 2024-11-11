package mongoDBTracerService

import (
	accessLogServicePort "app/port/accessLog"
	"context"
	"time"
)

type (
	DBQueryTracerService struct {
		AccessLogger accessLogServicePort.IAccessLogger
	}
)

func (this *DBQueryTracerService) Trace(collectionName string, label string, ctx context.Context) (end func(error)) {

	defer func() {

		if end == nil {

			end = func(error) {}
		}
	}()

	if this.AccessLogger == nil || !this.AccessLogger.IsLogging(ctx) {

		return
	}

	begin := time.Now()

	return func(err error) {

		dur := time.Since(begin)

		q := &MongoDBQueryDetail{
			Label:      "db_call",
			DBType:     "mongod",
			QueryType:  label,
			Collection: collectionName,
			DBTime:     float64(dur) / float64(time.Millisecond),
			dbTimeDur:  dur,
			Err:        err,
		}

		this.AccessLogger.PushTraceLogs(ctx, q)
	}
}
