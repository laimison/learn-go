//usr/bin/env go run "$0" "$@"; exit
//
// Try to run:
// ./command-line-flags.go -word=great -numb=10 -fork -word2=awesome
// ./command-line-flags.go

package main

import (
  "fmt"
  "flag"
)

var _ = fmt.Printf

func main() {
  // Basic flag declarations are available for string, integer, and boolean options. Here we declare a
  // string flag `word` with a default value `"foo"` and a short description. This `flag.String` function
  // returns a string pointer (not a string value); we'll see how to use this pointer below.
  wordPtr := flag.String("word", "foo", "a string")

  // It's also possible to declare an option that uses an existing var declared elsewhere in the program.
  // Note that we need to pass in a pointer to the flag declaration function.
  var word2 string
  flag.StringVar(&word2, "word2", "bar", "a string var")

  // `number` flag, using a similar approach
  numbPtr := flag.Int("number", 42, "an int")

  // `fork` flag, using a similar approach
  boolPtr := flag.Bool("fork", false, "a bool")

  // Once all flags are declared, call `flag.Parse()` to execute the command-line parsing.
  flag.Parse()

  // Note that we need to dereference the pointers with e.g. `*wordPtr` to get the actual option values.
  fmt.Println("word:", *wordPtr)
  fmt.Println("word2:", word2)
  fmt.Println("number:", *numbPtr)
  fmt.Println("fork:", *boolPtr)

  // fmt.Println("All arguments:", flag.Args())
}
