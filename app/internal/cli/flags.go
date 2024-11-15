package cli

import "flag"

var (
	trace_on *bool
)

func init() {

	trace_on = flag.Bool("trace", false, "display trace log in request log")

	flag.Parse()
}

func TraceLog() bool {

	return *trace_on
}
