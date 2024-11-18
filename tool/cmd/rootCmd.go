package cmd

import (
	"bch-tool/cmd/cache"
	"bch-tool/cmd/logCmd"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {

	RootCmd.AddCommand(cache.CacheBaseCommand)
	RootCmd.AddCommand(logCmd.LogBaseCmd)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {

	}
}
