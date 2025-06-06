package opLog

const (
	LOG_LEVEL_TRACE = "trace"
	LOG_LEVEL_DEBUG = "debug"

	LOG_CONTEXT_DEFAULT  = "default"
	LOG_CONTEXT_ADAPTIVE = "adaptive"
)

func empty_trace_func(err error) {}

func empty_push_cond_func(err error, msgIfErr string) {}

func empty_push_cond_with_messurement_func(msgIfNoErr string, err error, msgIfErr string) {}
