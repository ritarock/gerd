package cmd

import (
	"github.com/ritarock/gerd/internal/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "show created mermaid file",
	Long: `show mermaid file.
access to http://localhost:8080`,

	RunE: func(cmd *cobra.Command, args []string) error {
		server.Start()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
