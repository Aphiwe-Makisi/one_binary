package scripts

import "github.com/Aphiwe-Makisi/one_binary/internal/scripts/hello"

func init() {
	rootCmd.AddCommand(hello.NewCmd())
}
