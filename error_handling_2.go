//usr/bin/env go run "$0" "$@"; exit

package main

import (
  "fmt"
  "errors"
)

var _ = fmt.Printf

func Increment(n int) (int, error) {
  if n < 0 {
    // return error object
    return 1, errors.New("math: cannot process negative number")
  }
  return (n + 1), nil
}
func main() {
  // Change this number to negative to see the error
  // num := -1
  num := 5

  if inc, err := Increment(num); err != nil {
    fmt.Printf("Failed Number: %v, error message: %v\n", num, err)
  } else {
    fmt.Printf("Incremented Number: %v\n", inc)
  }
}
