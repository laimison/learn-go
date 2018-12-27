//usr/bin/env go run "$0" "$@"; exit
// can only send messages to the channel but cannot receive messages

package main

import (
  "fmt"
)

var _ = fmt.Printf

func main() {
 ch := make(chan string)

 go sc(ch)
 fmt.Println(<-ch)
}

func sc(ch chan<- string) {
 ch <- "hello"
}
