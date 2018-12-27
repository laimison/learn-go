//usr/bin/env go run "$0" "$@"; exit

package main

import (
  "fmt"
)

var _ = fmt.Printf

type person struct {
  name string
  age int
  gender string
}

func main() {
  //way 1: specifying attribute and value
  p := person{name: "Bob", age: 42, gender: "Male"}
  //way 2: specifying only value
  // person{"Bob", 42, "Male"}

  fmt.Println(p.name)
  fmt.Println(p.age)
  fmt.Println(p.gender)

  // You can also access attributes of a struct directly with its pointer
  pp := &person{name: "Bob", age: 42, gender: "Male"}

  fmt.Println(pp.name)
}
