package main

import (
	"errors"
	"fmt"
	"sort"
)

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

var inventory []Product

func main() {
	for {
		fmt.Println("\nInventory Management System")
		fmt.Println("1. Add Product")
		fmt.Println("2. Update Stock")
		fmt.Println("3. Search Product")
		fmt.Println("4. Display Inventory")
		fmt.Println("5. Sort Inventory")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addProduct()
		case 2:
			updateStock()
		case 3:
			searchProduct()
		case 4:
			displayInventory()
		case 5:
			sortInventory()
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addProduct() {
	var id, stock int
	var name string
	var price float64

	fmt.Print("Enter Product ID: ")
	fmt.Scanln(&id)
	fmt.Print("Enter Product Name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter Product Price: ")
	fmt.Scanln(&price)
	fmt.Print("Enter Product Stock: ")
	fmt.Scanln(&stock)

	if price <= 0 {
		fmt.Println("Error: Price must be greater than zero.")
		return
	}
	if stock < 0 {
		fmt.Println("Error: Stock cannot be negative.")
		return
	}

	if _, err := findProductByID(id); err == nil {
		fmt.Println("Error: Product ID already exists.")
		return
	}

	product := Product{
		ID:    id,
		Name:  name,
		Price: price,
		Stock: stock,
	}
	inventory = append(inventory, product)
	fmt.Println("Product added successfully.")
}
func updateStock() {
	var id, stock int
	fmt.Print("Enter Product ID: ")
	fmt.Scanln(&id)
	fmt.Print("Enter New Stock Quantity: ")
	fmt.Scanln(&stock)

	if stock < 0 {
		fmt.Println("Error: Stock cannot be negative.")
		return
	}

	product, err := findProductByID(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	product.Stock = stock
	fmt.Println("Stock updated successfully.")
}

func searchProduct() {
	fmt.Print("Search by 1. ID or 2. Name? Enter choice: ")
	var choice int
	fmt.Scanln(&choice)

	if choice == 1 {
		var id int
		fmt.Print("Enter Product ID: ")
		fmt.Scanln(&id)
		product, err := findProductByID(id)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Product Found: %+v\n", *product)
		}
	} else if choice == 2 {
		var name string
		fmt.Print("Enter Product Name: ")
		fmt.Scanln(&name)
		product, err := findProductByName(name)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Product Found: %+v\n", *product)
		}
	} else {
		fmt.Println("Invalid choice.")
	}
}
func displayInventory() {
	if len(inventory) == 0 {
		fmt.Println("No products in inventory.")
		return
	}
	fmt.Printf("%-10s %-20s %-10s %-10s\n", "ID", "Name", "Price", "Stock")
	for _, product := range inventory {
		fmt.Printf("%-10d %-20s %-10.2f %-10d\n", product.ID, product.Name, product.Price, product.Stock)
	}
}

func sortInventory() {
	fmt.Print("Sort by 1. Price or 2. Stock? Enter choice: ")
	var choice int
	fmt.Scanln(&choice)

	if choice == 1 {
		sort.Slice(inventory, func(i, j int) bool {
			return inventory[i].Price < inventory[j].Price
		})
		fmt.Println("Products sorted by price.")
	} else if choice == 2 {
		sort.Slice(inventory, func(i, j int) bool {
			return inventory[i].Stock < inventory[j].Stock
		})
		fmt.Println("Products sorted by stock.")
	} else {
		fmt.Println("Invalid choice.")
	}
}
func findProductByID(id int) (*Product, error) {
	for i := range inventory {
		if inventory[i].ID == id {
			return &inventory[i], nil
		}
	}
	return nil, errors.New("Product not found with this ID.")
}

func findProductByName(name string) (*Product, error) {
	for i := range inventory {
		if inventory[i].Name == name {
			return &inventory[i], nil
		}
	}
	return nil, errors.New("Product not found with this name.")
}
