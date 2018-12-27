//usr/bin/env go run "$0" "$@"; exit

package main

import (
  "fmt"
  "net/http"
)

var _ = fmt.Printf

func main(){
  resp, err := http.Get("http://example.com/")
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(resp)
}
