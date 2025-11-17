package scripts

import (
	"fmt"

	"github.com/Aphiwe-Makisi/one_binary/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "goscripts",
	Short: "A CLI exposing multiple Go scripts in a single binary",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return initConfig()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yml", "config file (yaml format)")
}

func initConfig() error {
	viper.SetConfigFile(cfgFile)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&config.C); err != nil {
		return err
	}

	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}
}
