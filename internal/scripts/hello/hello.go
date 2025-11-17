package hello

import (
	"fmt"

	"github.com/Aphiwe-Makisi/one_binary/internal/config"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:   "hello",
		Short: "Prints Hello, World!",
		Run: func(cmd *cobra.Command, args []string) {
			db := config.C.Database
			_ = db
			fmt.Printf("Hello, %s\n!", name)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "World", "Name to greet")

	return cmd
}
