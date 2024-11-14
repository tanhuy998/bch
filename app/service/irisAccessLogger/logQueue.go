package irisAccessLoggerService

import (
	accessLogServicePort "app/port/accessLog"
	"app/valueObject/log"
	"sync"
)

const (
	LOG_EVENT_THRESHOLD = 2
)

type (
	access_log_queue struct {
		sync.Mutex
		events       chan []interface{}
		end          chan struct{}
		logObj       *log.HTTPLogLine
		traceEnabled bool
	}
)

func (this *access_log_queue) Start() {

	this.events = make(chan []interface{}, LOG_EVENT_THRESHOLD)
	this.end = make(chan struct{})

	go func() {

		for newLines := range this.events {

			this._pushLogs(newLines...)
		}

		this.end <- struct{}{}
	}()
}

func (this *access_log_queue) _pushLogs(lines ...interface{}) {

	this.Lock()
	defer this.Unlock()

	logObj := this.logObj

	for _, l := range lines {

		if v, ok := l.(accessLogServicePort.IDBLogLine); ok {

			logObj.DBDuration += v.GetDBDuration()
		}
	}

	if len(logObj.TraceLogs) == 0 {

		logObj.TraceLogs = lines
		return
	}

	logObj.TraceLogs = append(logObj.TraceLogs, lines...)
}

func (this *access_log_queue) Stop() {

	close(this.events)
	this.Lock()
	defer this.Unlock()
	<-this.end
	close(this.end)
}

func (this *access_log_queue) Push(lines ...interface{}) {

	this.events <- lines
}
