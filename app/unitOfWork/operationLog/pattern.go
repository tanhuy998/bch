package opLog

import (
	"app/valueObject/log"
	"encoding/json"
)

type (
	inverse_mashaler logPattern

	logPattern struct {
		Operation  string `json:"operation,omitempty"`
		Service    string `json:"service,omitempty"`
		LogUnit    string `json:"log_unit,omitempty"`
		Message    string `json:"message,omitempty"`
		LogLevel   string `json:"level,omitempty"`
		ErrorMsg   string `json:"error_msg,omitempty"`
		LogContext string `json:"log_context,omitempty"`
		log.LogDurationInfo
	}
)

func NewLogPattern(level string, op string, msg string) *logPattern {

	return &logPattern{
		Operation: op,
		Message:   msg,
		LogLevel:  level,
	}
}

func newDefaultContextLogPattern(level, op, msg string) *logPattern {

	return &logPattern{
		Operation:  op,
		Message:    msg,
		LogLevel:   level,
		LogContext: LOG_CONTEXT_DEFAULT,
	}
}

func newAdaptiveContextLogPattern(level, op, msg string) *logPattern {

	return &logPattern{
		Operation:  op,
		Message:    msg,
		LogLevel:   level,
		LogContext: LOG_CONTEXT_ADAPTIVE,
	}
}

func newCustomLogPattern(base logPattern, custom interface{}) interface{} {

	baseBytes, err := json.Marshal(base)

	if err != nil {

		return nil
	}

	customBytes, err := json.Marshal(custom)

	if err != nil {

		return nil
	}

	var (
		baseMap   map[string]string
		customMap map[string]string
	)

	err = json.Unmarshal(baseBytes, &baseMap)

	if err != nil {

		return nil
	}

	json.Unmarshal(customBytes, &customMap)

	for key, v := range customMap {

		baseMap[key] = v
	}

	return baseMap
}
