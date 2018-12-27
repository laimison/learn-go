//usr/bin/env go run "$0" "$@"; exit

package main

import (
  "fmt"
)

var _ = fmt.Printf

func main() {
  // Array
  var a [2]int
  fmt.Println("Array a", a)

  // Multidimensional Array
  var d [2][3]int
  fmt.Println("Multidimensional array d", d)

  // Slices
  var b []int
  fmt.Println("Slice b", b)

  numbers := make([]int,5,10)
  fmt.Println("Slice numbers", numbers)

  numbers = append(numbers, 1, 2, 3, 4)
  fmt.Println("Slice numbers", numbers)

  var number([]int)
  number2 := make([]int, 10)
  copy(number2, number)
  fmt.Println("Slice number2", number2)

  // initialize a slice with 4 len and values
  number2 = []int{1,2,3,4}
  fmt.Println(numbers) // -> [1 2 3 4]
  // create sub slices
  slice1 := number2[2:]
  fmt.Println(slice1) // -> [3 4]
  slice2 := number2[:3]
  fmt.Println(slice2) // -> [1 2 3]
  slice3 := number2[1:4]
  fmt.Println(slice3) // -> [2 3 4]

  // Maps are a data type in Go, which maps keys to values. We can define a map using the following command
  var m map[string]int
  // adding key/value
  m['simplicity'] = 3
  // printing the values
  fmt.Println(m['clearity']) // -> 2
  fmt.Println(m['simplicity']) // -> 3
}
