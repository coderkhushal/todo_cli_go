/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"kbtodo/cmd/types"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// seetodosCmd represents the seetodos command
var seetodosCmd = &cobra.Command{
	Use:   "seetodos",
	Short: "gives the next todo timewise",
	Long:  `Gives the next todo in the line`,
	Run: func(cmd *cobra.Command, args []string) {

		_, err := os.Stat(`todos.json`)
		if os.IsNotExist(err) {

			fmt.Println(types.Red, "Todos Empty - Json file not found", err, types.Reset)
			return
		}
		file, err := os.ReadFile("./todos.json")
		if err != nil {
			fmt.Println(types.Red, "Something went wrong while reading file", types.Reset)
			return
		}
		var filecontent []types.TodoData

		json.Unmarshal(file, &filecontent)

		for i, value := range filecontent {
			fmt.Println(i+1, types.Yellow, strings.TrimSpace(value.Title), "", types.Reset, value.Time, "\n", types.Green, value.Description, "\n  Status:", value.Status, types.Reset)
		}

	},
}

func init() {
	rootCmd.AddCommand(seetodosCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// seetodosCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// seetodosCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
