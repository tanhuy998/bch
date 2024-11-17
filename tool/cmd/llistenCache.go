package cmd

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

const (
	CACHE_LOG_PORT = 3358
)

var (
	conn net.Conn
)

var listenCachCmd = &cobra.Command{
	Use: "cache",
	Run: func(cmd *cobra.Command, args []string) {

		ListenCache()
	},
}

func ListenCache() {

	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", CACHE_LOG_PORT))

	if err != nil {
		log.Println(err)
		return
	}

	go listenSIGINT()

	buffer := make([]byte, 30)

	for conn != nil {

		_, err := conn.Read(buffer)

		if err != nil {

			conn.Close()
			panic(err)
		}

		os.Stdout.Write(buffer)

		flush(buffer)
	}
}

func listenSIGINT() {

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	for range c {

		if conn == nil {

			return
		}

		conn.Close()
	}
}

func init() {

	rootCmd.AddCommand(listenCachCmd)
}

func flush(b []byte) {

	for i := range b {

		b[i] = 0
	}
}
