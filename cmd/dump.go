package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/sundowndev/go-search/engine"
)

func init() {
	// Register command
	rootCmd.AddCommand(dumpCmd)
}

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump database keys",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := engine.NewRedisClient(redisAddr, redisPort)
		if err != nil {
			fmt.Println("Failed to connect to database", redisAddr, redisPort)
			os.Exit(1)
		}

		fmt.Printf("Dumping 15 last keys...\n\n")

		keys, _ := client.GetAllKeys()

		for _, key := range keys {
			value, _ := client.GetKey(key)
			fmt.Println(key, value)
		}
	},
}
