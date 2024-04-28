package main

import (
	"fmt"
	"go-design-pattern/lesson/products"
	"time"
)

func main() {
	factory := products.Product{}

	product := factory.New()

	fmt.Println("My product was created at", product.CreatedAt.UTC())

	product2 := products.Product{
		ProductName: "Tomoya",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	fmt.Println("My product2 was created at", product2.CreatedAt.UTC())
}
