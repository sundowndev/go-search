package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/sundowndev/go-search/engine"
	"gopkg.in/yaml.v2"
)

type dumpResultFile struct {
	Path        string `yaml:"path"`
	Occurrences int    `yaml:"occurrences"`
}

type dumpResult struct {
	Word  string            `yaml:"word"`
	Files []*dumpResultFile `yaml:"files"`
}

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
		defer client.Close()

		var results []*dumpResult

		keys, err := client.GetAllKeys()
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		for _, key := range keys {
			value, err := client.GetKey(key)
			if err != nil {
				log.Fatalf("error: %v", err)
			}

			var files []*dumpResultFile

			for _, path := range value {
				files = append(files, &dumpResultFile{
					Path:        path,
					Occurrences: engine.CountWord("", key),
				})
			}

			results = append(results, &dumpResult{
				Word:  key,
				Files: files,
			})
		}

		data, err := yaml.Marshal(&results)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Printf("%s\n\n", string(data))
	},
}
