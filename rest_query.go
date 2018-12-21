//usr/bin/env go run "$0" "$@"; exit

package main

import (
  "fmt"
  "time"
  "net/http"
  "log"
  "io/ioutil"
  "encoding/json"
  _ "strconv"
)

func rest_query(method string, address string, user string, pass string) (string, string) {
  // Using http with timeout
  var client_timeout = &http.Client{
    Timeout: time.Second * 10,
  }

  // Doing query
  request, error := http.NewRequest(method, address, nil)

  // Basic HTTP authentication
  if len(user) >= 1 && len(pass) >= 1 {
    request.SetBasicAuth(user, pass)
  }

  // Getting the output
  response, error := client_timeout.Do(request)

  // Throw error
  if error != nil{
      log.Fatal(error)
  }

  // Write the output to memory
  bodyText, error := ioutil.ReadAll(response.Body)

  // Define JSON structure
  type MyStructure struct {
    id int
  }

  // Getting the output
  // Method 1
  var result MyStructure
  // Method 2
  // var result map[string]interface{}

  json.Unmarshal([]byte(bodyText), &result)

  // Method 1
  // my_variable := strconv.Itoa(result.id)
  // my_variable2 := strconv.Itoa(result.id)

  // Method 1 A
  // my_variable := ""
  // my_variable2 := ""

  // Method 2
  my_variable := string(bodyText)
  my_variable2 := ""

  return my_variable, my_variable2
}

func main() {
  // Json starting with [
  // fmt.Println(rest_query("GET", "https://jsonplaceholder.typicode.com/users", "", ""))

  // Json starting with {
  fmt.Println(rest_query("GET", "https://jsonplaceholder.typicode.com/users/1", "", ""))


}
