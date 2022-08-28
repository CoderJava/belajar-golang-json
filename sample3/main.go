package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func main() {
	reader, err := os.Open("assets/product.json")
	if err != nil {
		panic(err)
	}

	dataProduct := product{}

	// untuk mengubah json yang ada didalam file ke struct
	decoder := json.NewDecoder(reader)
	decoder.Decode(&dataProduct)

	fmt.Println(dataProduct)
}
