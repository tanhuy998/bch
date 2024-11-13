package log

import (
	libError "app/internal/lib/error"
	"time"
)

type (
	HTTPLogLine struct {
		*RequestExposedInfo
		HiddenInfo
		ResponseInfo
		IdentityInfo
		StartTime       time.Time     `json:"start_time"`
		RequestTime     float64       `json:"request_duration_ms"`
		DatabaseTime    float64       `json:"db_time_ms"`
		UseCase         string        `json:"usecase"`
		Error           string        `json:"error,omitempty"`
		Error_CallStack interface{}   `json:"error_call_stack,omitempty"`
		TraceLogs       []interface{} `json:"trace_logs"`
	}
)

func NewHTTPLogLine() *HTTPLogLine {

	ret := new(HTTPLogLine)

	ret.StartTime = time.Now()

	return ret
}

func (this *HTTPLogLine) WriteLogs(line ...interface{}) {

	this.TraceLogs = append(this.TraceLogs, line...)
}

func (this *HTTPLogLine) End() {

	this.RequestTime = float64(time.Since(this.StartTime)) / float64(time.Millisecond)
	this.DatabaseTime = float64(this.DBDuration) / float64(time.Millisecond)

	if this.Err != nil {

		this.Error = this.Err.Error()

		if v, ok := any(this.Err).(libError.ICallStackError); ok {

			this.Error_CallStack = v.CallStack()
		}
	}
}
