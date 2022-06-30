package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"testing"
)

var filepath string = "./messages.json"

func TestJsonRead(t *testing.T) {
	content := ReadJson(filepath)

	var testVar []uint8

	if reflect.TypeOf(content) != reflect.TypeOf(testVar) {
		t.Errorf("Error: ReadJson returned type %T, instead of %T", content, testVar)
	}  else {
		fmt.Println("Everything checked out for the test on Jsonread")
	}
}

var payload []Message
var m = Message{
	Episode: 100,
	Month:   "June",
	Title:   "Snazzy Title Here",
	Name:    "SD-1-1-11",
}


func TestJsonToMessage(t *testing.T) {
	content := ReadJson(filepath)

	err := json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatalf("%v", err)
	}

	payload = append(payload, m)

	if payload[len(payload)-1].Episode != m.Episode {
		t.Errorf("Messages are not being added as expected")
	} else {
		fmt.Println("Everything checked out for the test on JsonToMessage")
	}

}
