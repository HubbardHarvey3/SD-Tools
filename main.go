package main

import (
	"encoding/json"
	"strconv"
	"io/ioutil"
	"log"
	"os"
	"sort"
  "fmt"
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
  strInt, err := strconv.Atoi(os.Args[4]) 
  newMessage.Episode = int32(strInt) 

  fmt.Printf("%v\n", newMessage)



  // Read messages.json
  // Need to read from orderedMessages.json or overwrite messages.json
  content, err := ioutil.ReadFile("./messages.json")
  if err != nil {
    log.Fatal("Error when opening file: ", err)
  }
  var payload []Message
  err = json.Unmarshal(content, &payload)

  payload = append(payload, newMessage)

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

  output, err := json.Marshal(orderedPayload)
  ioutil.WriteFile("orderedMessages.json", output, os.ModePerm)

}
