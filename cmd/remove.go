package cmd

import (
	"encoding/json"
	"fmt"
	"kbtodo/cmd/types"
	"os"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "removes todo",
	Long:  `Updates todo by selecting todo`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Stat("./todos.json")
		if os.IsNotExist(err) {
			fmt.Println(types.Red, "File Does Not Exists", types.Reset)
			return
		}
		file, err := os.ReadFile("./todos.json")

		if err != nil {
			fmt.Println(types.Red, "Error while opening file", types.Reset)
			return
		}

		var filecontent []types.TodoData
		err = json.Unmarshal(file, &filecontent)

		if err != nil {
			fmt.Println(types.Red, "Error while parsing the file", types.Reset)
			return
		}
		var options []string
		for i, value := range filecontent {
			options = append(options, strconv.Itoa(i+1)+"."+value.Title)
		}
		prompt := &survey.Select{
			Message: "Choose Todo",
			Options: options,
		}
		var selectedOption string

		survey.AskOne(prompt, &selectedOption)
		numoption, err := strconv.Atoi(selectedOption[:1])
		if err != nil {
			fmt.Println(types.Red, "Some error occured", types.Reset)
			return
		}
		numoption--

		filecontent = append(filecontent[:numoption], filecontent[numoption+1:]...)
		updatedfilecontent, err := json.Marshal(filecontent)

		if err != nil {
			fmt.Println(types.Red, "Some error occured", types.Reset)
			return
		}
		err = os.WriteFile("./todos.json", updatedfilecontent, 0644)

		if err != nil {
			fmt.Println(types.Red, "Some error occured", types.Reset)
			return
		}
		fmt.Println(types.Green, "Removed Succesfully", types.Reset)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
