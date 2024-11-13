package log

import "time"

type (
	LogDurationInfo struct {
		Duration time.Duration `json:"-"`
		Time     float64       `json:"duration_ms,omitempty"`
	}
)

func (this *LogDurationInfo) SetDuration(dur time.Duration) {

	this.Duration = dur

	this.Time = float64(dur) / float64(time.Millisecond)
}
