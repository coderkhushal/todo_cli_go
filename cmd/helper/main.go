package helper

import (
	"encoding/json"
	"fmt"
	"kbtodo/cmd/types"
	"os"
	"time"
)

func GetCurrentDate() (string, error) {
	location, err := time.LoadLocation("EST")
	if err != nil {
		return "", err
	}

	t := time.Now().In(location).Format("2006-01-02")
	return t, nil
}

func GetCurrentTime() (string, error) {
	location, err := time.LoadLocation("EST")
	if err != nil {
		return "", err
	}

	t := time.Now().In(location).Format("15:04:05")
	return t, nil
}

func WriteToJson(a types.TodoData) {
	_, err := os.Stat("./todos.json")
	if os.IsNotExist(err) {
		fmt.Println(types.Red, "file does not exists", types.Reset)
	}
	file, err := os.ReadFile("./todos.json")
	if err != nil {
		fmt.Println(types.Red, "Something went wrong while Opening file", err, types.Reset)

	}
	var filecontent []types.TodoData

	err = json.Unmarshal(file, &filecontent)

	if err != nil {
		fmt.Println(types.Red, "Something went wrong while decoding the file", types.Reset)

	}
	filecontent = append(filecontent, a)

	updatedFileContent, err := json.Marshal(filecontent)

	if err != nil {
		fmt.Println(types.Red, "Something went wrong while marshaling file", types.Reset)

	}

	err = os.WriteFile("./todos.json", updatedFileContent, 0644)
	if err != nil {
		fmt.Println(types.Red, err, types.Reset)
	}
	fmt.Println(types.Green, "Updated Succesfully", types.Reset)
}
