package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var redisAddr string
var redisPort string

func init() {
	// Register flags
	rootCmd.PersistentFlags().StringVar(&redisAddr, "redis-addr", "localhost", "Redis server address")
	rootCmd.PersistentFlags().StringVar(&redisPort, "redis-port", "6379", "Redis server port")
}

var rootCmd = &cobra.Command{
	Use:     "gosearch [COMMANDS] [OPTIONS]",
	Short:   "A simple CLI search engine for your file system backed by Redis",
	Example: "gosearch index $(pwd)/fixtures",
}

// Execute is a function that executes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
