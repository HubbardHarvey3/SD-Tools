package main

import (
  "encoding/csv"
  "encoding/json"
  "errors"
  "fmt"
  "io/ioutil"
  "log"
  "os"
  "sort"
  "strconv"
)

// Reads the json at messages.json and passes it to the JsonToMessage function so it can be marshaled and added to.
func ReadJson(filepath string) []byte {
  content, err := ioutil.ReadFile(filepath)
  if err != nil {
    log.Fatal("Error when opening file: ", err)
  }

  return content
}

// Function calls ReadJson and then merges the new Messages from the
// function CSVtoJSON() and merge them with the newly converted
// []Message types from the messages.json file.
func JsonToMessage(csvMessages []Message, filepath string) ([]Message, error) {
  var originalPayload []Message

  content := ReadJson(filepath)

  err := json.Unmarshal(content, &originalPayload)
  if err != nil {
    log.Fatalf("%v", err)
  }

  for i := 0; i <= len(csvMessages)-1; i++ {
    for j := 0; j <= len(originalPayload)-1; j++ {
      if originalPayload[j].Episode == csvMessages[i].Episode {
        fmt.Printf("OriginalPayload: %v\n", originalPayload[j].Episode)
        fmt.Printf("CSVPayload: %v\n", csvMessages[i].Episode)
        return nil, errors.New("Episode Number " + fmt.Sprintf("%v", csvMessages[i].Episode) + " in the CSV exists already in the Messages.json\n")
      }
    }
  }

  for i := 0; i <= len(csvMessages)-1; i++ {
    originalPayload = append(originalPayload, csvMessages[i])
  }

  return originalPayload, nil
}

// Takes the an unsorted Slice of Message and sorts them based on the Episode
// number.  The resulting []Message will be sorted with the highest episode
// number at index 0 and the smallest episode number at the end of the array
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

// Reads the CSV file labeled example.csv found in the root of the dir
func ReadCSV(filepath string) [][]string {
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

// Takes the results from ReadCsv and converts the [][]string to
// []Message type
func CSVtoJSON(strArr [][]string) []Message {
  var messages []Message
  var message Message
  for i := 1; i <= len(strArr)-1; i++ {
    // Convert string to int64
    number, _ := strconv.ParseInt(strArr[i][0], 10, 64)
    message.Episode = number
    message.Name = strArr[i][1]
    message.Title = strArr[i][2]
    message.Month = strArr[i][3]
    messages = append(messages, message)
  }

  fmt.Println("Adding the following records:")
  for i := 0; i <= len(messages)-1; i++ {
    fmt.Printf(" Filename: %v | Broadcast Title: %v | Month: %v | Episode: %v \n", messages[i].Name, messages[i].Title, messages[i].Month, messages[i].Episode)
  }

  proceed()

  return messages
}

func proceed() {
  fmt.Println("Proceed with the update: Y (Anything other than 'Y' will abort)")
  var resp string
  fmt.Scanln(&resp)

  if resp == "Y" {
    fmt.Println("Continuing")
  } else {
    fmt.Println("Aborting")
    os.Exit(1)
  }
}
