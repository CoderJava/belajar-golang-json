package main

import (
	"encoding/json"
	"os"
)

type product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func main() {
	writer, err := os.Create("assets/new_product.json")
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(writer)

	dataProduct := product{
		Id:       "P002",
		Name:     "MacBook M3",
		ImageUrl: "https://example.com/images.png",
	}
	encoder.Encode(dataProduct)
}
