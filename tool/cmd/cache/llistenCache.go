package cache

import (
	"bch-tool/lib"
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/spf13/cobra"
)

const (
	CACHE_LOG_PORT = 3358
)

var listenCachCmd = &cobra.Command{
	Use:   "listen",
	Short: "Listen to host log events",
	Long:  `Listen to host log events.`,
	Run: func(cmd *cobra.Command, args []string) {

		ListenCache()
	},
}

func init() {

	CacheBaseCommand.AddCommand(listenCachCmd)
}

func ListenCache() {

	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", CACHE_LOG_PORT))

	if err != nil {
		log.Println(err)
		return
	}

	go lib.ListenSIGINT(func() {

		conn.Close()
	})

	buffer := make([]byte, 30)

	for conn != nil {

		_, err := conn.Read(buffer)

		switch {
		case errors.Is(err, net.ErrClosed):
			return
		case err != nil:
			conn.Close()
			panic(err)
		}

		os.Stdout.Write(buffer)

		flush(buffer)
	}
}

func flush(b []byte) {

	for i := range b {

		b[i] = 0
	}
}
