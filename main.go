package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Message struct {
	Name    string
	Title   string
	Month   string
	Episode int32
}

func main() {
	// gather params from user

	testSlice := ReadCsv("./example.csv")
	PrintCsv(testSlice)

	var newMessage Message

	if len(os.Args[1:]) != 1 {
		fmt.Println("No params provided")
	} else {
		newMessage.Name = os.Args[1]
		newMessage.Title = os.Args[2]
		newMessage.Month = os.Args[3]
		strInt, _ := strconv.Atoi(os.Args[4])
		newMessage.Episode = int32(strInt)

		payload := JsonToMessage(newMessage, "./messages.json")

		orderedPayload := SortJson(payload)

		output, _ := json.Marshal(orderedPayload)

		ioutil.WriteFile("messages.json", output, os.ModePerm)

	}

}
