//usr/bin/env go run "$0" "$@"; exit
//
// Defer is something that will always get executed at the end of a function.
//
// As you notice, there is a defer statement which will make the program execute the line at the end of the execution of the program.
// Defer can also be used when we need something to be executed at the end of the function, for example closing a file.

package main

import (
  "fmt"
)

var _ = fmt.Printf

func main() {
    f()
    fmt.Println("Returned normally from f.")
}

func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    fmt.Println("Calling g.")
    g(0)
    fmt.Println("Returned normally from g.")
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    g(i + 1)
}
