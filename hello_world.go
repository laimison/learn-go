//usr/bin/env go run "$0" "$@"; exit

package main

import (
  "fmt"
)

func example(my_variable string) string {
  return my_variable
}

func main() {
  fmt.Println(example("Hello World"))
}
