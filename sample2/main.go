package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonString := `{"street":"Tipar","country":"Indonesia","postal_code":"12345"}`
	jsonBytes := []byte(jsonString)

	var product map[string]interface{}

	// untuk mengubah json ke map
	if err := json.Unmarshal(jsonBytes, &product); err != nil {
		panic(err)
	}

	fmt.Println(product)
}
