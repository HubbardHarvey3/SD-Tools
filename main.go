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
