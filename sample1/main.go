package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	product := map[string]interface{}{
		"id":        "P001",
		"name":      "Macbook M1",
		"image_url": "https://example.com/images.png",
	}

	// Untuk mengubah map ke json
	bytes, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}
	fmt.Println(bytes)
	fmt.Println(string(bytes))
}
