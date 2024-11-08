package accessLogServicePort

import (
	"context"
)

type (
	IAccessLogger[DB_Monitor_T any] interface {
		Init(ctx context.Context)
		GetDBMonitor(ctx context.Context) *DB_Monitor_T
		SetDBMonitor(monitor *DB_Monitor_T, ctx context.Context)
		PushTraceLogs(ctx context.Context, line ...interface{})
		EndContext(ctx context.Context)

		AsyncTask(ctx context.Context, fn func())
	}
)
