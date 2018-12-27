//usr/bin/env go run "$0" "$@"; exit

// Typecasting - convert into another variable type
package main

import (
  "fmt"
  "os"
  "strconv"
)

var _ = fmt.Printf

func main() {
  // Convert float to integer
  a := 1.1
  a1 := int(a)
  fmt.Println(a1)

  // Convert string to integer
  b := "100"
  b1, err := strconv.Atoi(b)
  if err != nil {
    os.Exit(1)
  }
  fmt.Println(b1)
}
