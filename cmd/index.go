package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"

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
		client, err := engine.NewRedisClient(redisAddr, redisPort, "", 0)
		if err != nil {
			fmt.Println("Failed to connect to database", redisAddr, redisPort)
			os.Exit(1)
		}
		defer client.Close()

		path := args[0]

		fmt.Printf("Walking %v...\n", path)

		files := engine.GetFilesFromDir(path)

		var wg sync.WaitGroup
		wg.Add(len(files) - 1)

		for _, file := range files {
			// Open File
			f, err := ioutil.ReadFile(file)
			if err != nil {
				panic(err)
			}

			if engine.IsTextFile(f) == false {
				continue
			}

			content := string(f)

			go client.AddFile(file, content, &wg)

			fmt.Println("Successfully indexed file", file)
		}

		wg.Wait()
	},
}
