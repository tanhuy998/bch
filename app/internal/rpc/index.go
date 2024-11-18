package rpc

import (
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

func Listen() {

	err := rpc.RegisterName("LogTools", new(LogTools))

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
