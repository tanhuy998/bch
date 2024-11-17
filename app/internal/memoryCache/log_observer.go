package memoryCache

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

const (
	MEM_CACHE_PORT = 3358
)

var (
	tcp_server net.Listener
)

func init() {

	t, err := net.Listen("tcp", fmt.Sprintf(":%d", MEM_CACHE_PORT))

	if err != nil {

		log.Fatal("error while initializing memory cache server", err.Error())
	}

	tcp_server = t

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	go func() {

		for range c {

			tcp_server.Close()
			return
		}
	}()
}

func init() {

	go func() {

		for tcp_server != nil {

			conn, err := tcp_server.Accept()

			switch {
			case errors.Is(err, net.ErrClosed):
				return
			case err != nil:
				log.Println("memory cache connection error:", err.Error())
				continue
			}

			log_mgr.AddListener(time.Now().String(), conn)
		}
	}()
}
