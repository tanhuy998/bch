package log

import "time"

type (
	HiddenInfo struct {
		DBDuration time.Duration `json:"-"`
		Err        error         `json:"-"`
	}
)
