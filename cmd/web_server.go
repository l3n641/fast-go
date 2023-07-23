package cmd

import (
	"fast-go/bootstrap"
	"github.com/spf13/cobra"
)

var webServerCmd = &cobra.Command{
	Use:     "web-server",
	Short:   "启动web服务",
	Example: "fast-go web-server",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func run() {
	bootstrap.RunServer()
}
