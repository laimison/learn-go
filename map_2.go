//usr/bin/env go run "$0" "$@"; exit
package main

import "fmt"

func main() {
  elements := map[string]map[string]string{
    "Li": map[string]string{
      "name":"Lithium",
      "state":"solid",
    },
    "Ne":  map[string]string{
      "name":"Neon",
      "state":"gas",
    },
  }

  if e, ok := elements["Li"]; ok {
    // map[name:Lithium state:solid]
    fmt.Println(elements["Li"])

    // Lithium
    fmt.Println(elements["Li"]["name"])

    // solid
    fmt.Println(e["state"])
  }
}
