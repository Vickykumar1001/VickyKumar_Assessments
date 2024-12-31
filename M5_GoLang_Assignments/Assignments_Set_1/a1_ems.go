package main

import (
	"errors"
	"fmt"
	"strings"
)

type Employee struct {
	ID         int
	Name       string
	Age        int
	Department string
}

var employees []Employee

func main() {
	for {
		fmt.Println("\nEmployee Management System")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Search Employee")
		fmt.Println("3. List Employees by Department")
		fmt.Println("4. Count Employees by Department")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addEmployee()
		case 2:
			searchEmployee()
		case 3:
			listEmployeesByDepartment()
		case 4:
			countEmployeesByDepartment()
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func addEmployee() {
	var id, age int
	var name, department string

	fmt.Print("Enter Employee ID: ")
	fmt.Scanln(&id)
	fmt.Print("Enter Employee Name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter Employee Age: ")
	fmt.Scanln(&age)
	fmt.Print("Enter Employee Department: ")
	fmt.Scanln(&department)

	if _, err := findEmployeeByID(id); err == nil {
		fmt.Println("Error: Employee ID must be unique.")
		return
	}

	if age <= 18 {
		fmt.Println("Error: Age must be greater than 18.")
		return
	}

	employees = append(employees, Employee{
		ID:         id,
		Name:       name,
		Age:        age,
		Department: department,
	})

	fmt.Println("Employee added successfully.")
}

func searchEmployee() {
	fmt.Println("Search by: 1. ID, 2. Name")
	var searchType int
	fmt.Scanln(&searchType)

	if searchType == 1 {
		var id int
		fmt.Print("Enter Employee ID: ")
		fmt.Scanln(&id)

		employee, err := findEmployeeByID(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Employee found: %+v\n", employee)
	} else if searchType == 2 {
		var name string
		fmt.Print("Enter Employee Name: ")
		fmt.Scanln(&name)

		employee, err := findEmployeeByName(name)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Employee found: %+v\n", employee)
	} else {
		fmt.Println("Invalid search type.")
	}
}

func listEmployeesByDepartment() {
	var department string
	fmt.Print("Enter Department: ")
	fmt.Scanln(&department)

	fmt.Printf("Employees in %s:\n", department)
	found := false
	for _, employee := range employees {
		if strings.EqualFold(employee.Department, department) {
			fmt.Printf("%+v\n", employee)
			found = true
		}
	}

	if !found {
		fmt.Println("No employees found in this department.")
	}
}

func countEmployeesByDepartment() {
	var department string

	fmt.Print("Enter Department: ")
	fmt.Scanln(&department)

	count := 0
	for _, employee := range employees {
		if strings.EqualFold(employee.Department, department) {
			count++
		}
	}

	fmt.Printf("Total employees in %s: %d\n", department, count)
}

func findEmployeeByID(id int) (*Employee, error) {
	for _, employee := range employees {
		if employee.ID == id {
			return &employee, nil
		}
	}
	return nil, errors.New("Employee not found with this ID.")
}

func findEmployeeByName(name string) (*Employee, error) {
	for _, employee := range employees {
		if strings.EqualFold(employee.Name, name) {
			return &employee, nil
		}
	}
	return nil, errors.New("Employee not found with this name.")
}
