package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Message struct {
	Name    string
	Title   string
	Month   string
	Episode int64
}

func main() {
	// gather params from user

	// if len(os.Args[1:]) != 1 {
	// 	fmt.Println("No params provided")
	// } else {
	// newMessage.Name = os.Args[1]
	// newMessage.Title = os.Args[2]
	// newMessage.Month = os.Args[3]
	// strInt, _ := strconv.Atoi(os.Args[4])
	// newMessage.Episode = int64(strInt)

	csvSlice := ReadCSV("./example.csv")

	newPayloads := CSVtoJSON(csvSlice)

	originalPayload, err := JsonToMessage(newPayloads, "./messages.json")
	if err != nil {
		fmt.Println(err)
	} else {
		orderedPayload := SortJson(originalPayload)

		output, _ := json.Marshal(orderedPayload)

		ioutil.WriteFile("messages.json", output, os.ModePerm)
	}

}
