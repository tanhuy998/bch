package cmd

var (
	debug_mode bool
)

func ToggleDebugMode(state bool) {

	debug_mode = state
}

func IsDebugging() bool {

	return debug_mode
}
