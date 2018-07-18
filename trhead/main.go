package main

import (
  "fmt"
  "sync"
  "time"
  "net/http"
  "github.com/jinzhu/gorm"
  "io/ioutil"
  "encoding/json"
  "runtime"
  "os"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
  gorm.Model
  Code string
  Price uint
}

type ApiResults struct {
	Results []Person	`json:"results"`
  Info Info          `json:"info"`
}

type Person struct {
	Gender string	  `json:"gender"`
	Name User		    `json:"name"`
}

type Info struct {
  Seed string		  `json:"seed"`
  Results int		  `json:"results"`
  Page int		    `json:"page"`
  Version string		`json:"version"`
}
type User struct {
  Title string		`json:"title"`
  First string		`json:"first"`
  Last string		  `json:"last"`
}

type Request struct {
  Operation string      `json:"operation"`
  Key string            `json:"key"`
  Value string          `json:"value"`
  Info Info          `json:"info"`
}

func main() {
  runtime.GOMAXPROCS(2)

  var wg sync.WaitGroup
  wg.Add(2)

  t := time.Now()
  fmt.Println(t.Format(time.RFC850))
  fmt.Printf("start with thread -----\n")

	go func() {
    defer wg.Done()
    for i := 0; i < 50; i++ {
      data := getUser()
      if data == nil {
        fmt.Printf("is empty")
      }
    }
  }()

	go func() {
    defer wg.Done()
    for i := 0; i < 50; i++ {
      data := getUser()
      if data == nil {
        fmt.Printf("is empty")
      }
    }
  }()
  wg.Wait()
  fmt.Printf("End with thread-----\n")

  t = time.Now()
  fmt.Println(t.Format(time.RFC850))

  fmt.Printf("start without thread -----\n")
  for i := 0; i < 100; i++ {
    data := getUser()
    if data == nil {
      fmt.Printf("is empty")
    }
  }
  t = time.Now()
  fmt.Println(t.Format(time.RFC850))
  fmt.Printf("end all-----\n")

}


func getUser() *ApiResults{

    response, respErr := http.Get("https://randomuser.me/api/?inc=gender,name,na")
    data := &ApiResults{}
    if respErr != nil {
      fmt.Printf("%s", respErr)
      os.Exit(1)
    } else {
      defer response.Body.Close()
      contents, err := ioutil.ReadAll(response.Body)
      if err != nil {
          fmt.Printf("%s", err)
          os.Exit(1)
      }
      json.Unmarshal([]byte(contents), data)
      //fmt.Printf("%s%s\n", i, data.Results[0].Name.First)
    }
    return data
}

