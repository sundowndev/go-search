package cmd

import (
	"fmt"
	"io/ioutil"
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

		path := args[0]

		fmt.Printf("Walking %v...\n", path)

		for _, file := range engine.GetFilesFromDir(path) {
			// Open File
			f, err := ioutil.ReadFile(file)
			if err != nil {
				panic(err)
			}

			if engine.IsTextFile(f) == false {
				return
			}

			content := string(f)

			client.AddFile(file, content)

			fmt.Println("Successfully indexed file", file)
		}
	},
}
