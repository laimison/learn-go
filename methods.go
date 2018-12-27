//usr/bin/env go run "$0" "$@"; exit

package main

import (
  "fmt"
)

var _ = fmt.Printf

// struct defination
type person struct {
  name   string
  age    int
  gender string
}

// method defination
func (p *person) describe() {
  fmt.Printf("%v is %v years old.\n", p.name, p.age)
}
func (p *person) setAge(age int) {
  p.age = age
}

func (p person) setName(name string) {
  p.name = name
}

func main() {
  pp := &person{name: "Bob", age: 42, gender: "Male"}
  pp.describe()
  // => Bob is 42 years old

  pp.setAge(45)
  fmt.Println(pp.age)
  //=> 45

  // Note that in the above example the value of age is changed,
  // whereas the value of name is not changed because the method setName is of the receiver type whereas setAge is of type pointer.

  pp.setName("Hari")
  fmt.Println(pp.name)
  //=> Bob
}
