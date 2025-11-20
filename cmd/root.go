package cmd

import (
	"fmt"

	"github.com/Aphiwe-Makisi/one_binary/config"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfg config.Root
	env string
)

var rootCmd = &cobra.Command{
	Use:   "goscripts",
	Short: "A CLI exposing multiple Go scripts in a single binary",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&env, "environment", "e", "development", "Environment (development|production)")
	cobra.OnInitialize(initConfig)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warn().Msg("No configuration file found")
			return
		}
		log.Fatal().Err(err).Msg("failed to read configuration file")
	}

	if err := viper.UnmarshalKey(env, &cfg); err != nil {
		log.Fatal().Err(err).Msg("Fatal error config file")
	}

	log.Info().Msgf("Configuration loaded for environment: [%s]", env)
}

// initialise logging based on config
