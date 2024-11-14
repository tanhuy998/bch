package rpc

import (
	libCommon "app/internal/lib/common"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"os/signal"
)

const (
	ENV_TRACE_LOG = "TRACE_LOG"
	RPC_PORT      = "3328"
)

var (
	interupt bool
)

var (
	tcpListener net.Listener
)

func init() {

	t, err := net.Listen("tcp", fmt.Sprintf(`:%s`, RPC_PORT))

	if err != nil {

		log.Fatal("error occurs while initializing tcp connection for rpc server", err)
	}

	tcpListener = t

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	go func() {

		for range c {

			tcpListener.Close()
			interupt = true
			return
		}
	}()
}

type (
	RpcServer struct {
	}
)

func (this *RpcServer) ToggleTraceLog(request string, reply *string) error {

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

func Listen() {

	err := rpc.RegisterName("ToggleTraceLog", new(RpcServer))

	if err != nil {

		log.Fatal(err)
	}

	log.Println("rpc server listen on port 3328")

	for !interupt && tcpListener != nil {

		conn, err := tcpListener.Accept()

		if err != nil {

			log.Fatal("rpc error:", err)
		}

		go rpc.ServeConn(conn)
	}
}
