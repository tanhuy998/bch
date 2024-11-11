package accessLogServicePort

import (
	"context"
	"time"
)

type (
	IDBLogLine interface {
		GetDBDuration() time.Duration
	}

	IAccessLogger interface {
		Init(ctx context.Context)
		// GetDBMonitor(ctx context.Context) *DB_Monitor_T
		// SetDBMonitor(monitor *DB_Monitor_T, ctx context.Context)
		PushTraceLogs(ctx context.Context, line ...interface{})
		EndContext(ctx context.Context)
		IsLogging(ctx context.Context) bool
		PushError(context.Context, error)
		//AsyncTask(ctx context.Context, fn func())
	}
)
