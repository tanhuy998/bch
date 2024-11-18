package rpc

import (
	libCommon "app/internal/lib/common"
	"fmt"
	"log"
	"os"
)

type (
	LogTools struct {
	}
)

func (this *LogTools) ToggleTraceLog(request string, reply *string) error {

	if request != "on" && request != "off" {

		*reply = "invalid value of ToggleTraceLog"
		return nil
	}

	traceLogState := libCommon.Ternary(request == "on", "true", "false")

	err := os.Setenv(ENV_TRACE_LOG, traceLogState)

	if err != nil {

		log.Default().Println("rpc call RpcServer.ToggleTraceLog occurs error:", err)
	}

	*reply = libCommon.Ternary(err == nil, fmt.Sprintf(`trace log %s`, traceLogState), "an error occurs")

	return nil
}

func (this *LogTools) MemCacheStat(request string, reply *string) error {

	return nil
}
