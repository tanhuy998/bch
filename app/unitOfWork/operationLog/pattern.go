package opLog

import (
	"app/valueObject/log"
)

type (
	logPattern struct {
		Operation string `json:"operation,omitempty"`
		Message   string `json:"message,omitempty"`
		ErrorMsg  string `json:"error_msg,omitempty"`
		log.LogDurationInfo
	}
)
