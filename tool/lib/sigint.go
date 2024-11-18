package lib

import (
	"os"
	"os/signal"
)

var (
	sigintChan = make(chan os.Signal, 1)
)

func ListenSIGINT(fn func()) {

	signal.Notify(sigintChan, os.Interrupt)

	for range sigintChan {

		fn()
	}
}
