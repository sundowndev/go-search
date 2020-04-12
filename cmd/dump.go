package cmd

import (
	"fmt"
	"io/ioutil"
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

		// var results []*dumpResult
		var filesf []map[string]map[string]int

		files, err := client.GetAllKeys()
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		for _, file := range files {
			words, err := client.GetKey(file)
			if err != nil {
				log.Fatalf("error: %v", err)
			}

			// Open File
			f, err := ioutil.ReadFile(file)
			if err != nil {
				panic(err)
			}
			fileContent := string(f)

			for _, word := range words {
				filesf = append(filesf, map[string]map[string]int{
					word: map[string]int{
						file: engine.CountWord(fileContent, word),
					},
				})
			}
		}

		data, err := yaml.Marshal(&filesf)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Printf("%s\n\n", string(data))
	},
}
