//usr/bin/env go run "$0" "$@"; exit

package main

import (
  "fmt"
  "encoding/json"
)

var _ = fmt.Printf

func main(){
  mapA := map[string]int{"apple": 5, "lettuce": 7}
  mapB, _ := json.Marshal(mapA)
  fmt.Println(string(mapB))
}
