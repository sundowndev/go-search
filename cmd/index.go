package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/sundowndev/go-search/engine"
)

func init() {
	// Register command
	rootCmd.AddCommand(indexCmd)
}

var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Add files to database indexation",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := engine.NewRedisClient(redisAddr, redisPort)
		if err != nil {
			fmt.Println("Failed to connect to database", redisAddr, redisPort)
			os.Exit(1)
		}
		defer client.Close()

		filePath := args[0]

		fmt.Printf("Walking %v...\n", filePath)

		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

		client.AddFile(filePath, string(data))

		fmt.Println("Successfully indexed file", filePath)
	},
}
