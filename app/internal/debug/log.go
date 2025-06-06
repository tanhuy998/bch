package debug

import (
	"io"
	"log"
)

var (
	log_channels debug_log_channel_t
)

type (
	debug_log_channel_t struct {
		channel_list []*log.Logger
	}
)

func (this *debug_log_channel_t) NewChannel(writer io.Writer) {

	this.channel_list = append(this.channel_list, log.New(writer, "", 0))
}

func (this *debug_log_channel_t) Write(log interface{}) {

	for _, logger := range this.channel_list {

		logger.Print(log)
	}
}
