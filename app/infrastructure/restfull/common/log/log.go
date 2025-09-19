package log

import (
	"log"
	"os"
)

var (
	logger = log.New(os.Stderr, "[Restfull API]", log.Flags())
)

func Logger() *log.Logger {

	return logger
}
