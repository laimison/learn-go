//usr/bin/env go run "$0" "$@"; exit

package main

import (
  "fmt"
  "encoding/json"
)

var _ = fmt.Printf

type my_structure struct {
  PageNumber int `json:"page"`
  Fruits []string `json:"fruits"`
}

func main(){
  // While decoding the json byte using unmarshal,
  // the first argument is the json byte and the second argument is the address of the response type struct where
  // we want the json to be mapped to. Note that the json:”page” maps page key to PageNumber key in the struct.
  input := `{"page": 1, "fruits": ["apple", "peach"]}`
  output := my_structure{}
  json.Unmarshal([]byte(input), &output)
  fmt.Println(output.PageNumber)
  // => 1
}
