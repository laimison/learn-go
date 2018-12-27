//usr/bin/env go run "$0" "$@"; exit
// With this channel, we can send string type data. We can both send and receive data in this channel.
// The receiver Channels wait until the sender sends data to the channel.

package main

import (
  "fmt"
)

var _ = fmt.Printf

func main(){
  c := make(chan string)
  go func(){ c <- "hello" }()
  msg := <-c
  fmt.Println(msg)
}
//=>"hello"
