package main

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	// "log"
	"os"
	// "sort"
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

	var newMessage Message

	newMessage.Name = os.Args[1]
	newMessage.Title = os.Args[2]
	newMessage.Month = os.Args[3]
	strInt, _ := strconv.Atoi(os.Args[4])
	newMessage.Episode = int32(strInt)

	

	payload := JsonToMessage(newMessage)

	orderedPayload := SortJson(payload)

	output, _ := json.Marshal(orderedPayload)

	ioutil.WriteFile("messages.json", output, os.ModePerm)

}
