from book_management import add_book, view_books, search_book
from customer_management import add_customer, view_customers
from sale_management import sell_book, view_sales

books = []
customers = []
sales = []

def main():
    while True:
        print("\nWelcome to BookMart!")
        print("1. Book Management")
        print("2. Customer Management")
        print("3. Sales Management")
        print("4. Exit")
        choice = input("Enter your choice: ")

        if choice == "1":
            print("\nBook Management")
            print("1. Add Book")
            print("2. View Books")
            print("3. Search Book")
            sub_choice = input("Enter your choice: ")
            if sub_choice == "1":
                title = input("Title: ")
                author = input("Author: ")
                try:
                    price = float(input("Price: "))
                    quantity = int(input("Quantity: "))
                    if price <= 0 or quantity <= 0:
                        print("Error: Price and quantity must be positive numbers.")
                    else:
                        print(add_book(books, title, author, price, quantity))
                except ValueError:
                    print("Error: Price and quantity must be valid numbers.")
            elif sub_choice == "2":
                for book in view_books(books):
                    print(book)
            elif sub_choice == "3":
                title = input("Enter title to search: ")
                for book in search_book(books, title):
                    print(book)

        elif choice == "2": 
            print("\nCustomer Management")
            print("1. Add Customer")
            print("2. View Customers")
            sub_choice = input("Enter your choice: ")
            if sub_choice == "1":
                name = input("Name: ")
                email = input("Email: ")
                phone = input("Phone: ")
                print(add_customer(customers, name, email, phone))
            elif sub_choice == "2":
                for customer in view_customers(customers):
                    print(customer)

        elif choice == "3":
            print("\nSales Management")
            print("1. Sell Book ")
            print("2. View Sales")
            sub_choice = input("Enter your choice: ")
            if sub_choice == "1":
                customer_name = input("Customer Name: ")
                book_title = input("Book Title: ")
                try:
                    quantity = int(input("Quantity: "))
                    print(sell_book(books, sales, customer_name, book_title, quantity))
                except ValueError:
                    print("Error: Quantity must be a valid number.")
            elif sub_choice == "2":
                for sale in view_sales(sales):
                    print(sale)

        elif choice == "4":
            break

        else:
            print("Invalid choice..!! Please try again.")

if __name__ == "__main__":
    main()
