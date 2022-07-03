package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

func ReadJson(filepath string) []byte {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	return content
}

func JsonToMessage(m Message, filepath string) []Message {
	var payload []Message

	content := ReadJson(filepath)

	err := json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatalf("%v", err)
	}

	payload = append(payload, m)

	return payload
}

func SortJson(payload []Message) []Message {
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

func ReadCsv(filepath string) [][]string {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Unable to read input file "+filepath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filepath, err)

	}

	return records
}

func PrintCsv(csv [][]string) {
	for _, record := range csv {
		fmt.Println(record)
	}
}
