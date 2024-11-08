package log

import "time"

type (
	HTTPLogLine[DBMonitor_T any] struct {
		HttpVerb        string        `json:"http_verb"`
		Path            string        `json:"path"`
		ResponseStatus  int           `json:"response_status"`
		StartTime       time.Time     `json:"start_time"`
		SourceIP        string        `json:"source_ip"`
		UserAgent       string        `json:"user_agent"`
		RequestDuration float64       `json:"request_duration_ms"`
		UseCase         string        `json:"usecase"`
		DBMonitor       *DBMonitor_T  `json:"db_monitor,omitempty"`
		TraceLogs       []interface{} `json:"trace_logs"`
	}
)

func NewHTTPLogLine[DBMonitor_T any]() *HTTPLogLine[DBMonitor_T] {

	ret := new(HTTPLogLine[DBMonitor_T])

	ret.StartTime = time.Now()

	return ret
}

func (this *HTTPLogLine[DBMonitor_T]) SetDBMonitor(m *DBMonitor_T) {

	this.DBMonitor = m
}

func (this *HTTPLogLine[DBMonitor_T]) WriteLogs(line ...interface{}) {

	this.TraceLogs = append(this.TraceLogs, line...)
}

func (this *HTTPLogLine[DBMonitor_T]) End() {

	this.RequestDuration = float64(time.Since(this.StartTime)) / float64(time.Millisecond)
}
