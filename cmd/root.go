/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"os"

	"github.com/ritarock/gerd/internal/action"
	"github.com/spf13/cobra"
)

type Flag struct {
	User     string
	Password string
	Address  string
	DbName   string
}

var flag = &Flag{}

func makeFlags(cmd *cobra.Command) {
	user, _ := cmd.Flags().GetString("user")
	password, _ := cmd.Flags().GetString("password")
	address, _ := cmd.Flags().GetString("address")
	dbName, _ := cmd.Flags().GetString("db")
	flag.User = user
	flag.Password = password
	flag.Address = address
	flag.DbName = dbName
}

func (f *Flag) check() bool {
	if f.User == "" || f.Password == "" || f.Address == "" || f.DbName == "" {
		return false
	}
	return true
}

var rootCmd = &cobra.Command{
	Use:   "gerd",
	Short: "make entity relationship diagram",
	RunE: func(cmd *cobra.Command, args []string) error {
		makeFlags(cmd)
		if !flag.check() {
			return errors.New("Required flags are missing")
		}

		if err := action.Run(
			flag.User,
			flag.Password,
			flag.Address,
			flag.DbName,
		); err != nil {
			return err
		}
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP(
		"user",
		"u",
		"",
		"database user name (required)",
	)
	rootCmd.Flags().StringP(
		"password",
		"p",
		"",
		"database password (required)",
	)
	rootCmd.Flags().StringP(
		"address",
		"a",
		"",
		"connection address (required)",
	)
	rootCmd.Flags().StringP(
		"db",
		"d",
		"",
		"connection db name (required)",
	)
}
