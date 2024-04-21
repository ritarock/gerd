package cmd

import (
	"os"

	"github.com/ritarock/gerd/internal/mermaid"
	"github.com/spf13/cobra"

	database "github.com/ritarock/gerd/internal/db"
)

const FILE_NAME = "mermaid.md"

var rootCmd = &cobra.Command{
	Use:   "gerd",
	Short: "make entity relationship diagram",
	RunE: func(cmd *cobra.Command, args []string) error {
		db, _ := cmd.Flags().GetString("db")
		address, _ := cmd.Flags().GetString("address")
		user, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")

		dbx := database.NewDbx()
		if err := dbx.Connect(db, address, user, password); err != nil {
			return err
		}

		tables, err := dbx.GetTables()
		if err != nil {
			return err
		}

		mermaid.Create(FILE_NAME)
		for _, table := range tables {
			column, err := dbx.GetDescribe(table)
			if err != nil {
				return err
			}
			if err := mermaid.CreateTable(table, column, FILE_NAME); err != nil {
				return err
			}
		}

		for _, table := range tables {
			references, err := dbx.GetReferences(table)
			if err != nil {
				return err
			}
			if len(references) == 0 {
				continue
			}
			for _, reference := range references {
				if err := mermaid.CreateReference(table, reference, FILE_NAME); err != nil {
					return err
				}
			}
		}

		return nil
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
