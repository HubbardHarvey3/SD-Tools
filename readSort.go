package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"io/ioutil"
	"log"
)

func ReadJson() []byte {
	fmt.Printf("Test/n")
	// Read messages.json
	// Need to read from orderedMessages.json or overwrite messages.json
	content, err := ioutil.ReadFile("./messages.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	return content
}

func JsonToMessage(m Message) []Message {
	var payload []Message

	content := ReadJson()

	err := json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatalf("%v", err)
	}

	payload = append(payload, m)

	return payload
}

func SortJson(payload []Message) ([]Message) {
	var sortedOut []int

	for i := 0; i < len(payload); i++ {
		sortedOut = append(sortedOut, int(payload[i].Episode))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sortedOut)))

	var beginning int
	var end int
	beginning = sortedOut[0]
	end = sortedOut[len(sortedOut)-1]
	var orderedPayload []Message

	for i := beginning; i >= end; i-- {
		for v := range sortedOut {
			if i == int(payload[v].Episode) {
				orderedPayload = append(orderedPayload, payload[v])
			}
		}
	}

	return orderedPayload
}