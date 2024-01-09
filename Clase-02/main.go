package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func main() {

	filePath := "products.json"

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading the file: ", err)
		return
	}

	var products []Product

	err = json.Unmarshal(content, &products)

	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println("Products:")
	for _, p := range products {
		fmt.Printf("ID: %d, Name: %s, Quantity: %d, CodeValue: %s, IsPublished: %t, Expiration: %s, Price: %.2f\n",
			p.Id, p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price)
	}

}
