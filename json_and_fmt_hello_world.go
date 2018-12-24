//usr/bin/env go run "$0" "$@"; exit

// Unstructured JSON example
//
// As a general rule of thumb,
// if you can use structs to represent your JSON data, you should use them.
//
// The only good reason to use maps (for unstructured JSON) would be if it were not possible to use structs
// due to the uncertain nature of the keys or values in the data.

package main

import (
  "fmt"
  "encoding/json"
)

var _ = fmt.Printf

func main() {
  input := `{"some":"json"}`

  // All options should work
  // var myStoredVariable map[string]interface{}
  // myStoredVariable := map[string]interface{}{}
  // myStoredVariable := map[string]string{}
  var myStoredVariable map[string]interface{}

  json.Unmarshal([]byte(input), &myStoredVariable)

  fmt.Printf("%v\n", myStoredVariable) // is a generic placeholder. It will automatically convert your variable into a string with some default options. This is typically useful when printing primitives such as strings or numbers and you don’t need specific options.
  fmt.Printf("%#v\n", myStoredVariable) // prints your variable in Go syntax. This means you could copy the output and paste it into your code and it’ll syntactically correct. I find this most useful when working with structs and slices because it will print out the types and field names.
  fmt.Printf("%T\n", myStoredVariable) // prints your variable’s type. This is really useful for debugging if your data is passed as an interface{} and you want to see what its concrete type.
  fmt.Printf("%d\n", myStoredVariable) // prints an integer in base-10. You can do the same with %v but this is more explicit.
  fmt.Printf("%x\n", myStoredVariable) // and %X print an integer in base-16. One nice trick though is that you can pass in a byte slice and it’ll print each byte as a two-digit hex number.
  fmt.Printf("%f\n", myStoredVariable) // prints a floating point number without an exponent. You can do the same with %v but this becomes more useful when we add width and precision flags.
  fmt.Printf("%q\n", myStoredVariable) // prints a quoted string. This is useful when your data may have invisible characters (such as zero width space) because the quoted string will print them as escape sequences.
  fmt.Printf("%p\n", myStoredVariable) // prints a pointer address of your variable. This one is really useful when you’re debugging code and you want to check if different pointer variables reference the same data.

  fmt.Println(myStoredVariable["some"])
}
