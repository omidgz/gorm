package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"fmt"
)

type Product struct {
	gorm.Model
	Code        string
	Price       uint
	Description string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Starting")

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	// db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var products []Product
	db.Find(&products, &Product{})

	for _, p := range products {
		printProduct(&p)
	}

	/*
		db.First(&product, 1)                 // find product with integer primary key
		db.First(&product, "code = ?", "D42") // find product with code D42

		// Update - update product's price to 200
		db.Model(&product).Update("Price", 200)
		// Update - update multiple fields
		db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
		db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

		// Delete - delete product
		db.Delete(&product, 1)
	*/
}

func printProduct(p *Product) {
	fmt.Printf("ID: %d\tCode: %s\tDescription: %s\tPrice: %d\n", p.ID, p.Code, p.Description, p.Price)
}
