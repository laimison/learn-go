//usr/bin/env go run "$0" "$@"; exit

package main

import (
  "fmt"
)

var _ = fmt.Printf

func increment(i *int) {
  *i++
}
func main() {
  i := 10
  increment(&i)
  fmt.Println(i)
}
