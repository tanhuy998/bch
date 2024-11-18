package logCmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

const (
	HOST_PORT = 3328
)

var LogBaseCmd = &cobra.Command{
	Use:   "log",
	Short: "host log tools",
	Run: func(cmd *cobra.Command, args []string) {

		err := handleFlags(cmd)

		if err != nil {

			log.Fatal(err)
		}
	},
}

func init() {

	LogBaseCmd.PersistentFlags().String("trace", "on", "toggle trace log on host server")
}

func handleFlags(cmd *cobra.Command) error {

	traceFlag := cmd.Flag("trace").Value.String()

	switch {
	case traceFlag == "":
		traceFlag = "on"
		fallthrough
	case traceFlag == "on", traceFlag == "off":
		return toggleTraceLog(traceFlag)
	default:
		return fmt.Errorf("invalid value of log --trace flag")
	}
}
