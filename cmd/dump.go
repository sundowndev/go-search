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

func init() {
	// Register command
	rootCmd.AddCommand(dumpCmd)
}

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump database keys",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := engine.NewRedisClient(redisAddr, redisPort, "", 0)
		if err != nil {
			fmt.Println("Failed to connect to database", redisAddr, redisPort)
			os.Exit(1)
		}
		defer client.Close()

		var results = make(map[string]map[string]int)

		files, err := engine.GetFiles(client)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		for _, file := range files {
			// Open File
			f, err := ioutil.ReadFile(file)
			if err != nil {
				panic(err)
			}
			fileContent := string(f)

			for w, s := range engine.Scan(fileContent) {
				if results[w] == nil {
					results[w] = make(map[string]int)
				}

				results[w][file] = s
			}
		}

		data, err := yaml.Marshal(&results)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Printf("%s\n\n", string(data))
	},
}
