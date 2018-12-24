//usr/bin/env go run "$0" "$@"; exit
package main

import (
	"fmt"
	"encoding/json"
)

var _ = fmt.Printf

type PublicKey struct {
	Id int
	People string
  // [] tells that it's json array, remove if you don't have array in front of groups:
	Groups []struct {
		Name string
	}
}

func main() {
	input := []byte(`[
		{"id": 1, "people": "Tom", "groups": [{"name": "groupA"}]},
		{"id": 2, "people": "Ben", "groups": [{"name": "groupB"}]},
		{"id": 3, "people": "Dan", "groups": [{"name": "groupA"}, {"name": "groupC"}]}
		]`)

	var output []PublicKey
	json.Unmarshal([]byte(input), &output)

	for key, value := range output {
		fmt.Printf("Id: %v\n", key)
		fmt.Printf("People: %v\n", value.People)

		// fmt.Printf("Groups: %v\n", value.Groups[0].Name)
    for key_groups, value_groups := range value.Groups {
		  fmt.Printf("  Group %v: %v\n", key_groups, value_groups.Name)
		}
	  fmt.Println()
  }
}
