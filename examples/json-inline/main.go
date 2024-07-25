package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	A `json:",inline"`
}

type A struct {
	X int `json:"X"`
	Y int `json:"Y"`
}

func main() {
	data := []byte(`{
		"X": 1,
		"Y": 2
	}`)

	var foo Foo
	_ = json.Unmarshal(data, &foo)

	fmt.Printf("%+v\n", foo.A)
	fmt.Printf("%+v\n", foo.X)
	fmt.Printf("%+v\n", foo.Y)
}
