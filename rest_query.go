//usr/bin/env go run "$0" "$@"; exit

package main

import (
  "fmt"
  "time"
  "net/http"
  "log"
  "io/ioutil"
  "encoding/json"
)

func rest_query(method string, address string, user string, pass string) string {
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

  // Getting the output
  var result map[string]interface{}
  json.Unmarshal([]byte(bodyText), &result)

  content := string(bodyText)
  return content
}

func main() {
  // Json starting with [
  fmt.Println(rest_query("GET", "https://jsonplaceholder.typicode.com/posts", "", ""))

  // Json starting with {
  fmt.Println(rest_query("GET", "https://jsonplaceholder.typicode.com/todos/1", "", ""))
}
