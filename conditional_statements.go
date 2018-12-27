//usr/bin/env go run "$0" "$@"; exit

package main

import (
  "fmt"
)

var _ = fmt.Printf

func main() {
  // if else
  if num := 9; num < 0 {
    fmt.Println(num, "is negative")
  } else if num < 10 {
    fmt.Println(num, "has 1 digit")
  } else {
    fmt.Println(num, "has multiple digits")
  }

  // switch case
  i := 2
  switch i {
  case 1:
    fmt.Println("one")
  case 2:
    fmt.Println("two")
  default:
    fmt.Println("none")
  }

  // loop
  sum := 0
  for i := 0; i < 10; i++ {
    sum += i
  }
  fmt.Println(sum)

  // infinite loop
  // for {
  // }

}
