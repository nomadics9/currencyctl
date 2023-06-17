package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
  "time"
)


type Currencylayer struct {
	Success   bool   `json:"success"`
	Terms     string `json:"terms"`
	Privacy   string `json:"privacy"`
	Timestamp int    `json:"timestamp"`
	Source    string `json:"source"`
	Quotes    struct {
		Usdkwd float64 `json:"USDKWD"`
	} `json:"quotes"`
}


func main()  {
   url := fmt.Sprintf("http://apilayer.net/api/live?access_key=YOUR_KEY_HEREcurrencies=KWD&source=USD&format=1")

  // request
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    log.Fatal("NewRequest: ", err)
    return
  }
  
  // client
  client := &http.Client{}

  //response
  resp, err := client.Do(req)
  if err != nil {
    log.Fatal("Do: ", err)
    return
  }

  //closing
  defer resp.Body.Close()

  //fill
  var record Currencylayer

  //json decode
  if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
    log.Println(err)
  }
 
  //replace your path here $HOME not working for me
  f, err := os.Create("/home/nomad/go/src/github.com/nomadics9/currconvert/api/price.txt")
  if err != nil {
    log.Fatal(err)
  }

  defer f.Close()
  

  rate := fmt.Sprintf("%v", record.Quotes.Usdkwd)
  currentTime := time.Now() 

  myslice := []string {rate ,"\n","", "\n" ,currentTime.Format("15:04 02-01-2006") ,"\n"}
  result := strings.Join(myslice, "")
   


  _, err2 := f.WriteString(result)
  if err2 != nil {
    log.Fatal(err2)
  }
  fmt.Println("USD/KWD Updated")
  fmt.Println(currentTime.Format("15:04 02-01-2006"))


  //formula

   // var rate, usd, answer float64
   // rate = record.Quotes.Usdkwd
  // fmt.Println("💵USD/KWD💴")
  // fmt.Scanln(&usd)
  // answer = rate * usd
  // fmt.Println(answer)
 
}


