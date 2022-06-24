package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

type Message struct {
  Name string
  Title string
  Month string
  Episode int32
}

type Entry struct {
  Key string
  value *Message
}

type byPriority []Entry

func (d byPriority) Len() int {
  return len(d)
}
func (d byPriority) Less(i, j int) bool {
  return d[i].value.Episode < d[j].value.Episode
}
func (d byPriority) Swap(i, j int) {
  d[i], d[j] = d[j], d[i]
}

func printSorted(detail map[string]*Message) {
  // Copy entries into a slice.
  slice := make(byPriority, 0, len(detail))
  for key, value := range detail {
      slice = append(slice, Entry{key, value})
  }

  // Sort the slice.
  sort.Sort(slice)

  // Iterate and print the entries in sorted order.
  for _, entry := range slice {
      fmt.Printf("%s : %v\n", entry.Key, entry.value)
  }
}

func main()  {
  newMessage := Message{"SD-100-100-100", "Cool long title here", "Febarch", 581}

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
  // fmt.Printf("%s\n",content)
  var payload []Message
  err = json.Unmarshal(content, &payload)
  payload = append(payload, newMessage)
  fmt.Printf("%T", payload)

  var sortedOut map[string]*Message
  // for i := range payload {
  //   fmt.Printf("Name is : %v\n", payload[i].Episode)
  //   sortedOut = append(sortedOut, )
  // }
  
  
  printSorted(sortedOut)

  


}