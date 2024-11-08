package irisAccessLoggerService

import (
	"app/valueObject/log"
	"context"
	"encoding/json"
	"sync"

	stdLog "log"

	"github.com/kataras/iris/v12"
)

const (
	CTX_LOG_KEY = "custom_access_log"
)

type (
	access_log_queue[DB_T any] struct {
		sync.WaitGroup
		sync.Mutex
		logObj *log.HTTPLogLine[DB_T]
	}
)

type (
	IrisAccessLoggerService[DB_T any] struct {
		LogChannel *stdLog.Logger
	}
)

func (this *IrisAccessLoggerService[DB_T]) resolveContext(ctx context.Context) iris.Context {

	c, ok := ctx.(iris.Context)

	if !ok {

		panic("the given context is not type of iris.Context")
	}

	return c
}

func (this *IrisAccessLoggerService[DB_T]) getQueue(ctx context.Context) *access_log_queue[DB_T] {

	c := this.resolveContext(ctx)

	accessLogObj := c.Value(CTX_LOG_KEY)

	if accessLogObj == nil {

		return nil
	}

	if ret, ok := accessLogObj.(*access_log_queue[DB_T]); ok {

		return ret
	}

	panic("IrisAccessLoggerService error: could not resolve access log queue from the context")
}

func (this *IrisAccessLoggerService[DB_T]) getLogObject(ctx context.Context) *log.HTTPLogLine[DB_T] {

	c := this.resolveContext(ctx)

	accessLogObj := c.Value(CTX_LOG_KEY)

	if accessLogObj == nil {

		return nil
	}

	if ret, ok := accessLogObj.(*access_log_queue[DB_T]); ok {

		return ret.logObj
	}

	panic("IrisAccessLoggerService error: could not resolve access log object from the context")
}

func (this *IrisAccessLoggerService[DB_T]) Init(ctx context.Context) {

	c := this.resolveContext(ctx)

	accessLogObj := c.Value(CTX_LOG_KEY)

	if accessLogObj != nil {

		panic("IrisAccessLoggerService error: access logger initialzation at wrong place")
	}

	payload := access_log_queue[DB_T]{
		logObj: log.NewHTTPLogLine[DB_T](),
	}

	c.Values().Set(CTX_LOG_KEY, &payload)
}

func (this *IrisAccessLoggerService[DB_T]) GetDBMonitor(ctx context.Context) *DB_T {

	logObj := this.getLogObject(ctx)

	if logObj == nil {

		this.Init(ctx)
	}

	logObj = this.getLogObject(ctx)

	return logObj.DBMonitor
}

func (this *IrisAccessLoggerService[DB_T]) SetDBMonitor(monitor *DB_T, ctx context.Context) {

	logObj := this.getLogObject(ctx)

	if logObj == nil {

		this.Init(ctx)
	}

	logObj = this.getLogObject(ctx)

	logObj.DBMonitor = monitor
}

func (this *IrisAccessLoggerService[DB_T]) PushTraceLogs(ctx context.Context, lines ...interface{}) {

	logObj := this.getLogObject(ctx)

	if logObj == nil {

		this.Init(ctx)
	}

	logObj = this.getLogObject(ctx)
	queue := this.getQueue(ctx)

	go func() {

		queue.Lock()
		queue.WaitGroup.Add(1)
		defer queue.WaitGroup.Done()
		defer queue.Unlock()

		if logObj.TraceLogs == nil {

			logObj.TraceLogs = lines

			return
		}

		logObj.TraceLogs = append(logObj.TraceLogs, lines...)
	}()
}

func (this *IrisAccessLoggerService[DB_T]) EndContext(ctx context.Context) {

	if this.LogChannel == nil {

		this.LogChannel = stdLog.Default()
	}

	queue := this.getQueue(ctx)

	queue.Wait()

	this.assignLogObject(ctx)

	logObj := this.getLogObject(ctx)

	raw, _ := json.MarshalIndent(logObj, "", "\t")

	this.LogChannel.Println(string(raw))
}

func (this *IrisAccessLoggerService[DB_T]) assignLogObject(ctx context.Context) {

	c := this.resolveContext(ctx)
	logObj := this.getLogObject(ctx)

	logObj.End()

	logObj.HttpVerb = c.Method()
	logObj.Path = c.Path()
	logObj.SourceIP = c.GetHeader("X-Real-IP")
	logObj.UserAgent = c.Request().UserAgent()
	logObj.ResponseStatus = c.GetStatusCode()
}

func (this *IrisAccessLoggerService[DB_T]) AsyncTask(ctx context.Context, toDo func()) {

	queue := this.getQueue(ctx)

	if toDo == nil {

		return
	}

	queue.Add(1)

	go func() {

		queue.Lock()
		defer queue.Unlock()

		toDo()

		queue.Done()
	}()
}
