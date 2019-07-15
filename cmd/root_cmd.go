package cmd

import (
	"log"
	"os"

	"net-go-api/config"

	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use: "example",
	Run: run,
}

// RootCommand will setup and return the root command
func RootCommand() *cobra.Command {
	rootCmd.PersistentFlags().StringP("config", "c", "", "the config file to use")
	rootCmd.Flags().IntP("port", "p", 0, "the port to use")

	return &rootCmd
}

func run(cmd *cobra.Command, args []string) {
	conf, err := config.LoadConfig(cmd)
	if err != nil {
		log.Fatal("Failed to load config: " + err.Error())
	}

	logger, err := config.ConfigureLogging(&conf.LogConfig)
	if err != nil {
		log.Fatal("Failed to configure logging: " + err.Error())
	}

	server := api.NewAPI(logger, config)
	logger.Infof("Starting up server on port %d", config.Port)
	if err := server.Start(); err != nil {
		logger.WithError(err).Error("Error while running server")
		os.Exit(1)
	}
}
