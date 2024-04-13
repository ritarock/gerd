package cmd

import (
	"os"

	"github.com/ritarock/gerd/internal/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gerd",
	Short: "make entity relationship diagram",
	RunE: func(cmd *cobra.Command, args []string) error {
		db, _ := cmd.Flags().GetString("db")
		address, _ := cmd.Flags().GetString("address")
		user, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")

		return action.Run(db, address, user, password)
	},
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("db", "d", "", "connection db name (required)")
	rootCmd.Flags().StringP("address", "a", "", "connection address (required)")
	rootCmd.Flags().StringP("user", "u", "", "database user name (required)")
	rootCmd.Flags().StringP("password", "p", "", "database password (required)")

	rootCmd.MarkFlagRequired("db")
	rootCmd.MarkFlagRequired("address")
	rootCmd.MarkFlagRequired("user")
	rootCmd.MarkFlagRequired("password")
}
