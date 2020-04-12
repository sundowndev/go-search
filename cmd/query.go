package cmd

import (
	"fmt"
	"log"
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
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := engine.NewRedisClient(redisAddr, redisPort)
		if err != nil {
			fmt.Println("Failed to connect to database", redisAddr, redisPort)
			os.Exit(1)
		}
		defer client.Close()

		word := args[0]

		fmt.Printf("Querying index for \"%s\":\n\n", word)

		results, err := client.GetAllKeys()
		if err != nil {
			log.Fatal(err)
		}

		var queryResults []*engine.QueryResult

		for _, file := range results {
			score := client.GetScore(file, word)

			if score == 0 {
				continue
			}

			queryResults = append(queryResults, &engine.QueryResult{
				File:       file,
				Count:      score,
				FirstMatch: engine.GetFirstMatchLine("", word),
			})
		}

		engine.ShowResults(queryResults)
	},
}
