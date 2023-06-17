package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"

)

func main() {


    // Your Path here
    content, err := ioutil.ReadFile("/home/nomad/go/src/github.com/nomadics9/currconvert/api/price.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    // Convert content to string
    data := string(content)

    // Split the content by newline character to get individual lines
    lines := strings.Split(data, "\n")

    // Parse each line to float64
    for _, line := range lines {
        // Trim any leading or trailing whitespaces
        // line = strings.TrimSpace(line)

        // Skip empty lines
        if line == "" {
          break
         }

        // Parse the line to float64
        rate, err := strconv.ParseFloat(line, 64)
        if err != nil {
         fmt.Println("Error parsing line:", err)
           continue 
        }

        // Use the parsed float64 value
        // fmt.Println(value)
  //formula

    var usd, answer float64
   fmt.Println("💵USD/KWD💴")
   fmt.Scanln(&usd)
   answer = rate * usd
   fmt.Println(answer)
    }
}
