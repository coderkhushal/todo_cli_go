/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"kbtodo/cmd/types"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Updates todo status to be done",
	Long:  `Updates the todo status to be done`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Stat("./todos.json")
		if os.IsNotExist(err) {
			fmt.Println(types.Red, "Todos not found - json file not found error", types.Reset)
			return

		}
		file, err := os.ReadFile("./todos.json")
		if err != nil {
			fmt.Println(types.Red, "Error opening todos.json file ", types.Reset)
			return
		}

		var filecontent []types.TodoData
		err = json.Unmarshal(file, &filecontent)

		if err != nil {
			fmt.Println(types.Red, "Some error while parsing json file", types.Reset)
			return
		}
		fmt.Println(types.Yellow, "Enter Done Todo Number", types.Reset)
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		if err != nil || input == "" {
			fmt.Println(types.Red, "Error : enter some input ", types.Reset)
			return
		}

		n, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil || n > len(filecontent) {
			fmt.Println(types.Red, "Todo not found , Please enter valid todonumber", types.Reset)
			return
		}
		n--

		filecontent[n].Status = true

		updatedfilecontent, err := json.Marshal(filecontent)

		if err != nil {
			fmt.Println(types.Red, "Some error while writing the file", types.Reset)
			return
		}
		err = os.WriteFile("./todos.json", updatedfilecontent, 0644)

		if err != nil {
			fmt.Println(types.Red, "Some error while writing the file", types.Reset)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
