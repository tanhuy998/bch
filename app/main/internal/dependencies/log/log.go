package log

import (
	"log"
	"os"
)

var (
	logger = log.New(os.Stderr, "[Main]", log.Flags())
)

func Main() *log.Logger {

	return logger
}
