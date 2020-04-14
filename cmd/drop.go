package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/sundowndev/go-search/engine"
)

func init() {
	// Register command
	rootCmd.AddCommand(dropCmd)
}

var dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop all indexes",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := engine.NewRedisClient(redisAddr, redisPort, redisPassword, redisDB)
		if err != nil {
			fmt.Println("Failed to connect to database", redisAddr, redisPort)
			os.Exit(1)
		}
		defer client.Close()

		fmt.Println("Dropping database...")

		client.FlushAll()

		fmt.Println("Done.")
	},
}
