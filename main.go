package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "os"
  "sort"
)

type Message struct {
  Name    string
  Title   string
  Month   string
  Episode int32
}


func main() {
  newMessage := Message{"SD-100-100-100", "Cool long title here", "Febarch", 581}
  anotherMessage := Message{"SD-900-900-900", "Slightly Different Title", "Juntember", 582}
  // b, err := json.Marshal(newMessage)
  // if err != nil {
  // 	fmt.Print(err)
  // }
  // fmt.Printf("%s\n", b)

  // Read messages.json
  content, err := ioutil.ReadFile("./messages.json")
  if err != nil {
    log.Fatal("Error when opening file: ", err)
  }
  var payload []Message
  err = json.Unmarshal(content, &payload)
  payload = append(payload, newMessage)
  payload = append(payload, anotherMessage)
  fmt.Printf("%T", payload)

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
