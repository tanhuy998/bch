package debug

import (
	"app/internal/internal/cmd"
	"log"
)

const (
	DEBUG_CONTEXT_KEY = "debug"
)

func init() {

	log_channels.channel_list = make([]*log.Logger, 0)
}

func GetDebugLogChannelKey() interface{} {

	return DEBUG_CONTEXT_KEY
}

func IsDebugging() bool {

	return cmd.IsDebugging()
}

func GetDebugChannels() *debug_log_channel_t {

	return &log_channels
}
