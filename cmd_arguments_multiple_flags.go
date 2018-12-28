//usr/bin/env go run "$0" "$@"; exit

package main

import (
  "fmt"
  "flag"
)

var _ = fmt.Printf

type arrayFlags []string

func (i *arrayFlags) String() string {
  return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
  *i = append(*i, value)
  return nil
}

var ListFlags arrayFlags
var word string

func main() {
  // Var, StringVar, Int, Bool and others can be used
  flag.Var(&ListFlags, "list1", "Some description for this param.")
  flag.StringVar(&word, "word", "bar", "a string var")

  // flag.Set is just simulation of script.go -list1 value1 -list1 value2
  flag.Set("list1", "value1")
  flag.Set("list1", "value2")
  flag.Set("list1", "value3")
  // simulation ends

  flag.Parse()

  // check out flags
  fmt.Println(ListFlags)
  fmt.Println(word)
}
