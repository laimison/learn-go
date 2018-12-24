//usr/bin/env go run "$0" "$@"; exit

package main

import (
  "fmt"
  "encoding/json"
)

var _ = fmt.Printf

func example() (string, string) {
  type Bird struct {
    Species string
    Description string
  }

  birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
  var bird Bird
  json.Unmarshal([]byte(birdJson), &bird)

  my_variable := bird.Species
  my_variable2 := bird.Description

  return my_variable, my_variable2
}

func main() {
  my_example, my_example2 := example()
  fmt.Println(my_example, "|", my_example2)

}
