/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"kbtodo/cmd/helper"
	"kbtodo/cmd/types"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// addtodoCmd represents the addtodo command
var addtodoCmd = &cobra.Command{
	Use:   "addtodo",
	Short: "Adds a todo",
	Long: `Adds a todo , time is set to EST 
	--date yyyy-mm-dd ( default today)
	--time mm:ss 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		dateFlag, _ := cmd.Flags().GetString("date")
		timeFlag, _ := cmd.Flags().GetString("time")
		var data types.TodoData
		if dateFlag == "" {
			currentdate, err := helper.GetCurrentDate()
			if err != nil {
				data.Date = ""

			} else {
				data.Date = currentdate
			}

			fmt.Println(types.Yellow, "Date : Today", types.Reset)
		} else {
			data.Date = dateFlag
		}

		if timeFlag == "" {
			// currenttime, err := helper.GetCurrentTime()

			// if err != nil {
			// 	fmt.Println(types.Red, "something went wrong", types.Reset)
			// 	data.Time = ""

			// } else {
			// 	data.Time = currenttime
			// }

			// fmt.Println(types.Yellow, "Date : Today", types.Reset)
			fmt.Println(types.Red, "Please enter End Time ", types.Reset)
			return
		} else {

			data.Time = timeFlag
		}

		//taking input

		reader := bufio.NewReader(os.Stdin)
		fmt.Println(types.Green, "Enter Title", types.Reset)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(types.Red, "Input error", types.Reset)
			return
		}
		data.Title = strings.TrimSpace(input)
		fmt.Println(types.Green, "Enter description", types.Reset)
		input, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(types.Red, "Input error", types.Reset)
			return
		}
		data.Description = strings.TrimSpace(input)
		helper.WriteToJson(data)
		//Printing Done
		fmt.Println(types.Yellow, "Title: ", types.Green, data.Title, types.Reset)
		fmt.Println(types.Yellow, "Description: ", types.Green, data.Description, types.Reset)
		fmt.Println(types.Yellow, "Status: ", types.Green, data.Title, types.Reset)
		fmt.Println(types.Yellow, "Date: ", types.Green, data.Date, types.Reset)
		fmt.Println(types.Yellow, "Time: ", types.Green, data.Time, types.Reset)
	},
}

func init() {
	rootCmd.AddCommand(addtodoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	addtodoCmd.PersistentFlags().String("date", "", "Add date in yyyy-mm-dd format ( default today)")
	addtodoCmd.PersistentFlags().String("time", "", "Required flag , give time in hh:mm eg- 15:30")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addtodoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
