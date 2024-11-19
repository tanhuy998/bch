package opLog

import (
	"app/valueObject/log"
)

type (
	logPattern struct {
		Operation string `json:"operation,omitempty"`
		Service   string `json:"service,omitempty"`
		Message   string `json:"message,omitempty"`
		LogLevel  string `json:"level,omitempty"`
		ErrorMsg  string `json:"error_msg,omitempty"`
		log.LogDurationInfo
	}
)
