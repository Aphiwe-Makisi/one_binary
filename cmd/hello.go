package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var name string

func init() {
	helloCmd.Flags().StringVarP(&name, "name", "n", "World", "Name to greet")
	rootCmd.AddCommand(helloCmd)
}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints Hello, World!",
	Run: func(cmd *cobra.Command, args []string) {

		db := cfg.Database
		fmt.Printf("Connecting to database %s at %s:%d with user %s\n", db.Name, db.Host, db.Port, db.Username)

		fmt.Printf("Hello, %s\n!", name)
	},
}
