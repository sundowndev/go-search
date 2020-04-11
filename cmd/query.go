package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/sundowndev/go-search/engine"
)

func init() {
	// Register command
	rootCmd.AddCommand(queryCmd)
}

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Run a query against the database",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := engine.NewRedisClient(redisAddr, redisPort)
		if err != nil {
			fmt.Println("Failed to connect to database", redisAddr, redisPort)
			os.Exit(1)
		}

		print("query")
	},
}
