package main

import (
	"fmt"

	json2 "example.com/json-not-escape-html/json"
)

func main() {
	data := map[string]interface{}{
		"a": "<>&",
	}

	fmt.Println(data)

	for _, escapeHTML := range []bool{true, false} {
		result, err := json2.MarshalJSON(data, escapeHTML)
		fmt.Println(string(result), err)
	}
}
