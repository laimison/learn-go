//usr/bin/env go run "$0" "$@"; exit
package main

import "fmt"
import "encoding/json"

type PublicKey struct {
	Id int
	Key string
	Test []interface{}
}

func main() {
	keysBody := []byte(`[
		{"id": 1, "key": "-", "test": [{"name": "test1"}]},
		{"id": 2, "key": "-", "test": [{"name": "test1"}]},
		{"id": 3, "key": "-", "test": [{"name": "test1"}]}
		]`)
	var keys []PublicKey
	json.Unmarshal(keysBody, &keys)

	fmt.Printf("%v\n", keys)
}
