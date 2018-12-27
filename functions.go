//usr/bin/env go run "$0" "$@"; exit

package main

import (
  "fmt"
)

var _ = fmt.Printf

func add(a int, b int) (int, string) {
  c := a + b
  return c, "second-return-value"
}
func main() {
  fmt.Println(add(2, 1))
}
