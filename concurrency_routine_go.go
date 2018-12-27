//usr/bin/env go run "$0" "$@"; exit
// Go routines are the function which can run in parallel or concurrently with another function. Creating a Go routine is very simple.
// Simply by adding a keyword Go in front of a function, we can make it execute in parallel. Go routines are very lightweight,
// so we can create thousands of them.

package main

import (
  "fmt"
  "time"
)

var _ = fmt.Printf

func main() {
  go c()
  fmt.Println("I am main")
  time.Sleep(time.Second * 2)
}
func c() {
  // Not working in new version of Go?
  time.Sleep(time.Second * 2)
  fmt.Println("I am concurrent")
}
//=> I am main
//=> I am concurrent
