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
	rootCmd.AddCommand(queryCmd)
}

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Run a query against the database",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := engine.NewRedisClient(redisAddr, redisPort, "", 0)
		if err != nil {
			fmt.Println("Failed to connect to database", redisAddr, redisPort)
			os.Exit(1)
		}
		defer client.Close()

		word := args[0]

		fmt.Printf("Querying index for \"%s\":\n\n", word)

		files, err := client.GetFiles()
		if err != nil {
			log.Fatal(err)
		}

		var queryResults []*engine.QueryResult

		for _, file := range files {
			score := client.GetWordScoreFromFile(file, word)

			// If score is 0 then ignore it
			if score == 0 {
				continue
			}

			// Open File
			f, err := ioutil.ReadFile(file)
			if err != nil {
				panic(err)
			}
			text := string(f)

			queryResults = append(queryResults, &engine.QueryResult{
				File:       file,
				Score:      score,
				FirstMatch: engine.GetFirstMatchingLine(text, word),
			})
		}

		engine.ShowResults(engine.SortResultsByScore(queryResults))
	},
}
