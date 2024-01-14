package storage

import (
	"clase-02/internal"
	"encoding/json"
	"fmt"
	"os"
)

type StorageProductJSON struct {
	filePath string
}

func NewStorageProductJSON(filePath string) *StorageProductJSON {

	return &StorageProductJSON{
		filePath: filePath,
	}
}

type ProductJSON struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func (s *StorageProductJSON) ReadAll() (products []internal.Product, err error) {

	filePath := s.filePath

	content, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println("Error reading the file: ", err)
		return products, err
	}

	err = json.Unmarshal(content, &products)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return products, err
	}

	fmt.Println("Products from JSONFile have been readed into memory")

	return products, nil
}

// Write all replaces everything in the JSON db given
// a slice of Products

func (s *StorageProductJSON) WriteAll(products []internal.Product) (err error) {
	filePath := s.filePath

	// Convert products slice to JSON format
	jsonData, err := json.MarshalIndent(products, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling products to JSON:", err)
		return err
	}

	// Write JSON data to the file
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		return err
	}

	fmt.Println("Products have been written to JSON file")

	return nil
}
