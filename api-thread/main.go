package main

import (
  "fmt"
  "net/http"
  "github.com/jinzhu/gorm"
    "io/ioutil"
    "encoding/json"
    "os"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
  gorm.Model
  Code string
  Price uint
}

type ApiResults struct {
	Results []Person 	`json:"results"`
    Info Info          `json:"info"`
}

type Person struct {
	Gender string 	`json:"gender"`
	Name User 		`json:"name"`
}

type Info struct {
  Seed string 		`json:"seed"`
  Results int 		`json:"results"`
  Page int 		`json:"page"`
  Version string 		`json:"version"`
}
type User struct {
  Title string 		`json:"title"`
  First string 		`json:"first"`
  Last string 		`json:"last"`
}

type Request struct {
    Operation string      `json:"operation"`
    Key string            `json:"key"`
    Value string          `json:"value"`
    Info Info          `json:"info"`
}


//nos vamos a consultar una api de datos en internet y la llenamos en la tabla
//de datos de prueba
//https://randomuser.me/api/
func main() {
    response, err := http.Get("https://randomuser.me/api/?inc=gender,name,na")

    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
	userString := string(contents)
	//userString = string(`{"operation": "get", "key": "example", "info": {"seed":"xxx"}}`)
	//userString = string(`{"results":[{"gender":"male","name":{"title":"mr","first":"dylan","last":"welch"}}],"info":{"seed":"9fda9ddf1d5c2d92","results":1,"page":1,"version":"1.2"}}`)
	//userString = string(`{"operation": "get", "key": "example"}`)
        //fmt.Printf("%s\n", userString)
	data := &ApiResults{}
	//data := &Request{}
	json.Unmarshal([]byte(userString), data)


	fmt.Printf("%s\n", string(userString))
        //fmt.Printf("%s\n", data.results[0].name.first)
        fmt.Printf("%+v\n", data)


    }


  /*db, err := gorm.Open("sqlite3", "test.db")
  if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()




  // Migrate the schema
  /*db.AutoMigrate(&Product{})

  // Create
  db.Create(&Product{Code: "L1212", Price: 1000})

  // Read
  var product Product
  db.First(&product, 1) // find product with id 1
  db.First(&product, "code = ?", "L1212") // find product with code l1212

  // Update - update product's price to 2000
  db.Model(&product).Update("Price", 2000)

  // Delete - delete product
  db.Delete(&product)*/


}


