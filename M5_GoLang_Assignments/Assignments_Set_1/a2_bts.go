package main

import (
	"errors"
	"fmt"
)

type Account struct {
	ID                 int
	Name               string
	Balance            float64
	TransactionHistory []string
}

var accounts []Account

const (
	CreateAccountOption = 1
	DepositOption       = 2
	WithdrawOption      = 3
	ViewBalanceOption   = 4
	ViewHistoryOption   = 5
	ExitOption          = 6
)

func main() {
	for {
		fmt.Println("\nBank Transaction System")
		fmt.Println("1. Create Account")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. View Balance")
		fmt.Println("5. View Transaction History")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case CreateAccountOption:
			createAccount()
		case DepositOption:
			deposit()
		case WithdrawOption:
			withdraw()
		case ViewBalanceOption:
			viewBalance()
		case ViewHistoryOption:
			viewTransactionHistory()
		case ExitOption:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func createAccount() {
	var id int
	var name string
	fmt.Print("Enter New Account ID: ")
	fmt.Scanln(&id)
	fmt.Print("Enter Account Holder Name: ")
	fmt.Scanln(&name)

	if _, err := findAccountByID(id); err == nil {
		fmt.Println("Error: Account ID already exists.")
		return
	}
	newAccount := Account{
		ID:                 id,
		Name:               name,
		Balance:            0.0,
		TransactionHistory: []string{},
	}
	accounts = append(accounts, newAccount)
	fmt.Println("Account created successfully.")
}

func deposit() {
	var id int
	var amount float64
	fmt.Print("Enter Account ID: ")
	fmt.Scanln(&id)
	fmt.Print("Enter Deposit Amount: ")
	fmt.Scanln(&amount)

	if amount <= 0 {
		fmt.Println("Error: Deposit amount must be greater than zero.")
		return
	}
	account, err := findAccountByID(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	account.Balance += amount
	account.TransactionHistory = append(account.TransactionHistory, fmt.Sprintf("Deposited: %.2f", amount))
	fmt.Println("Deposit successful.")
}

func withdraw() {
	var id int
	var amount float64
	fmt.Print("Enter Account ID: ")
	fmt.Scanln(&id)
	fmt.Print("Enter Withdraw Amount: ")
	fmt.Scanln(&amount)

	if amount <= 0 {
		fmt.Println("Error: Withdraw amount must be greater than zero.")
		return
	}
	account, err := findAccountByID(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	if account.Balance < amount {
		fmt.Println("Error: Insufficient balance.")
		return
	}
	account.Balance -= amount
	account.TransactionHistory = append(account.TransactionHistory, fmt.Sprintf("Withdrew: %.2f", amount))
	fmt.Println("Withdrawal successful.")
}

func viewBalance() {
	var id int
	fmt.Print("Enter Account ID: ")
	fmt.Scanln(&id)

	account, err := findAccountByID(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Account Balance: %.2f\n", account.Balance)
}

func viewTransactionHistory() {
	var id int
	fmt.Print("Enter Account ID: ")
	fmt.Scanln(&id)

	account, err := findAccountByID(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Transaction History:")
	if len(account.TransactionHistory) == 0 {
		fmt.Println("No transactions found.")
		return
	}
	for _, transaction := range account.TransactionHistory {
		fmt.Println(transaction)
	}
}

func findAccountByID(id int) (*Account, error) {
	for i := range accounts {
		if accounts[i].ID == id {
			return &accounts[i], nil
		}
	}
	return nil, errors.New("Account not found with this ID.")
}
