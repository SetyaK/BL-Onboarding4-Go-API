package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

type Product struct {
	ProductID   int64
	Name        string
	Description string
	Stock       int
}

func TestGetProduct(t *testing.T) {
	resp, err := http.Get("http://localhost:4567/product/1")
	if err != nil {
		t.Fatal(err)
	}
	product := new(Product)
	json.NewDecoder(resp.Body).Decode(&product)
	fmt.Println(product)
}
