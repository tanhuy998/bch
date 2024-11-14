package irisAccessLoggerService

import (
	"app/cli"
	"app/internal/bootstrap"
	libCommon "app/internal/lib/common"
	"app/valueObject/log"
	"context"
	"encoding/json"
	"os"

	stdLog "log"

	"github.com/kataras/iris/v12"
)

const (
	ENV_TRACE_LOG = bootstrap.ENV_TRACE_LOG
	CTX_LOG_KEY   = "custom_access_log"
)

func init() {

	os.Setenv(
		ENV_TRACE_LOG,
		libCommon.Ternary(
			cli.TraceLog(), "true", os.Getenv(ENV_TRACE_LOG),
		),
	)
}

type (
	IrisAccessLoggerService struct {
		LogChannel *stdLog.Logger
	}
)

func (this *IrisAccessLoggerService) resolveContext(ctx context.Context) iris.Context {

	c, ok := ctx.(iris.Context)

	if !ok {

		panic("the given context is not type of iris.Context")
	}

	return c
}

func (this *IrisAccessLoggerService) getQueue(ctx context.Context) *access_log_queue {

	//c := this.resolveContext(ctx)

	if ctx == nil {

		return nil
	}

	accessLogObj := ctx.Value(CTX_LOG_KEY)

	if accessLogObj == nil {

		return nil
	}

	if ret, ok := accessLogObj.(*access_log_queue); ok {

		return ret
	}

	panic("IrisAccessLoggerService error: could not resolve access log queue from the context")
}

func (this *IrisAccessLoggerService) getLogObject(ctx context.Context) *log.HTTPLogLine {

	//c := this.resolveContext(ctx)

	accessLogObj := ctx.Value(CTX_LOG_KEY)

	if accessLogObj == nil {

		return nil
	}

	if ret, ok := accessLogObj.(*access_log_queue); ok {

		return ret.logObj
	}

	panic("IrisAccessLoggerService error: could not resolve access log object from the context")
}

func (this *IrisAccessLoggerService) Init(ctx context.Context) {

	c := this.resolveContext(ctx)

	accessLogObj := c.Value(CTX_LOG_KEY)

	if accessLogObj != nil {

		panic("IrisAccessLoggerService error: access logger initialzation at wrong place")
	}

	logQueue := &access_log_queue{
		logObj:       log.NewHTTPLogLine(),
		traceEnabled: os.Getenv(ENV_TRACE_LOG) == "true",
	}

	c.Values().Set(CTX_LOG_KEY, logQueue)

	logQueue.Start()
}

func (this *IrisAccessLoggerService) PushTraceLogs(ctx context.Context, lines ...interface{}) {

	if ctx == nil || !this.IsTraceLogging(ctx) {

		return
	}

	go func() {

		if !this.IsLogging(ctx) {

			return
		}

		queue := this.getQueue(ctx)

		if queue == nil {

			return
		}

		queue.Push(lines...)
	}()
}

func (this *IrisAccessLoggerService) EndContext(ctx context.Context) {

	this.assignLogObject(ctx)

	go func() {

		queue := this.getQueue(ctx)

		if queue == nil {

			return
		}

		queue.Stop()

		logObj := queue.logObj

		raw, _ := json.MarshalIndent(logObj, "", "\t")

		this.LogChannel.Println(string(raw))
	}()
}

func (this *IrisAccessLoggerService) assignLogObject(ctx context.Context) {

	//c := this.resolveContext(ctx)

	c, ok := ctx.(iris.Context)

	if !ok {

		return
	}

	queue := this.getQueue(ctx)

	if queue == nil {

		return
	}

	logObj := queue.logObj

	logObj.End()

	logObj.RequestExposedInfo = new(log.RequestExposedInfo)

	logObj.RequestType = "rest"
	logObj.IsSecure = c.IsSSL()
	logObj.Protocol = libCommon.Ternary(c.IsHTTP2(), "http/2", "http/1")
	logObj.Verb = c.Method()
	logObj.Path = c.Path()
	logObj.SourceIP = c.GetHeader("X-Real-IP")
	logObj.UserAgent = c.Request().UserAgent()
	logObj.ResponseStatus = c.GetStatusCode()
}
func (this *IrisAccessLoggerService) HasError(ctx context.Context) bool {

	return this.GetError(ctx) != nil
}

func (this *IrisAccessLoggerService) GetError(ctx context.Context) error {

	queue := this.getQueue(ctx)

	if queue == nil {

		return nil
	}

	return queue.logObj.Err
}

func (this *IrisAccessLoggerService) PushError(ctx context.Context, err error) {

	queue := this.getQueue(ctx)

	if queue == nil {

		return
	}

	queue.logObj.Err = err
}

func (this *IrisAccessLoggerService) IsLogging(ctx context.Context) bool {

	return this.getQueue(ctx) != nil
}

func (this *IrisAccessLoggerService) IsTraceLogging(ctx context.Context) bool {

	return this.getQueue(ctx).traceEnabled
}

func (this *IrisAccessLoggerService) WriteMessage(ctx context.Context, msg string) {

	queue := this.getQueue(ctx)

	if queue == nil {

		return
	}

	queue.logObj.Message = msg
}
